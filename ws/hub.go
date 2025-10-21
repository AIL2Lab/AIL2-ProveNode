package ws

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"

	"AIL2-ProveNode/db"
	"AIL2-ProveNode/ail2"
	"AIL2-ProveNode/log"
	"AIL2-ProveNode/types"
)

type Hub struct {
	wg      sync.WaitGroup
	wsConns sync.Map
	do      *delayOffline
	open    atomic.Bool
}

type cachedOfflineItem struct {
	disconnectTime time.Time
	stakingType    types.StakingType
}

type delayOfflineChanInfo struct {
	machine        types.MachineKey
	disconnectTime time.Time
	stakingType    types.StakingType
}

type delayOffline struct {
	connect   chan delayOfflineChanInfo
	diconnect chan delayOfflineChanInfo
	elements  map[types.MachineKey]cachedOfflineItem
	wg        sync.WaitGroup
	done      chan bool
	notifyApi string
}

func InitHub(ctx context.Context, napi string) (*Hub, error) {
	hub := &Hub{
		wg:      sync.WaitGroup{},
		wsConns: sync.Map{},
		do: &delayOffline{
			connect:   make(chan delayOfflineChanInfo),
			diconnect: make(chan delayOfflineChanInfo),
			elements:  make(map[types.MachineKey]cachedOfflineItem),
			wg:        sync.WaitGroup{},
			done:      make(chan bool),
			notifyApi: napi,
		},
	}
	if err := db.MDB.ReadDelayOffline(
		ctx,
		func(mk types.MachineKey, t time.Time, st types.StakingType) {
			hub.do.elements[mk] = cachedOfflineItem{
				disconnectTime: t,
				stakingType:    st,
			}
		},
	); err != nil {
		return nil, err
	}
	go hub.do.HandleDelayOffline()
	hub.open.Store(true)
	return hub, nil
}

func (h *Hub) Close() {
	h.open.Store(false)
	h.wsConns.Range(func(key, value any) bool {
		if conn, ok := key.(*Client); ok {
			conn.WriteEnvelope(envelope{t: websocket.CloseMessage, msg: []byte("server exiting")})
			// conn.WriteEnvelope(envelope{t: websocket.CloseMessage, msg: websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")})
			conn.conn.Close()
		}
		return true
	})
	log.Log.Println("Shutdownd all websocket connections")
}

func (h *Hub) closed() bool {
	return !h.open.Load()
}

func (h *Hub) Wait() {
	h.wg.Wait()
	log.Log.Println("All websocket routine exiting")
	h.do.done <- true
	h.do.wg.Wait()
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	db.MDB.WriteAllDelayOffline(ctx, time.Now())
}

func (h *Hub) SendUnregisterNotify(machine types.MachineKey, stakingType types.StakingType) {
	h.wsConns.Range(func(key, value any) bool {
		if client, ok := key.(*Client); ok {
			if client.MachineKey == machine && client.StakingType == stakingType {
				notify := types.WsNotifyMessage{
					Unregister: types.WsUnregisterNotify{
						Message: "machine unregistered, notify from http api",
					},
				}
				jsonBody, err := json.Marshal(notify)
				if err != nil {
					log.Log.WithFields(logrus.Fields{
						"machine": machine,
					}).Errorf("send unregister notify message failed %v", err)
				} else {
					client.WriteResponse(&types.WsResponse{
						WsHeader: types.WsHeader{
							Version:   0,
							Timestamp: time.Now().Unix(),
							Id:        0,
							Type:      uint32(types.WsMtNotify),
							PubKey:    []byte(""),
							Sign:      []byte(""),
						},
						Code:    0,
						Message: "notify",
						Body:    jsonBody,
					})
					log.Log.WithFields(logrus.Fields{
						"machine": machine,
					}).Info("begin to send unregister notify message")
				}
				return false
			}
		}
		return true
	})
}

func (do *delayOffline) HandleDelayOffline() {
	ticker := time.NewTicker(10 * time.Second)
	defer func() {
		ticker.Stop()
		close(do.connect)
		close(do.diconnect)
	}()
	for {
		select {
		case item := <-do.connect:
			delete(do.elements, item.machine)
		case item := <-do.diconnect:
			do.elements[item.machine] = cachedOfflineItem{
				disconnectTime: item.disconnectTime,
				stakingType:    item.stakingType,
			}
		case <-ticker.C:
			expired := time.Now().Add(-5 * time.Minute)
			for machine, cdi := range do.elements {
				if cdi.disconnectTime.Before(expired) {
					go do.Offline(delayOfflineChanInfo{
						machine:        machine,
						disconnectTime: cdi.disconnectTime,
						stakingType:    cdi.stakingType,
					})
					delete(do.elements, machine)
				}
			}
		case <-do.done:
			return
		}
	}
}

func (do *delayOffline) Offline(info delayOfflineChanInfo) {
	do.wg.Add(1)
	defer do.wg.Done()

	const maxRetries = 3
	retries := 0
	for retries < maxRetries {
		ctx1, cancel1 := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel1()
		if hash, err := dbc.DbcChain.Report(
			ctx1,
			types.MachineOffline,
			info.stakingType,
			info.machine.Project,
			info.machine.MachineId,
		); err != nil {
			log.Log.WithFields(logrus.Fields{
				"machine": info.machine,
			}).Errorf(
				"machine offline in chain contract failed with hash %v because of %v",
				hash,
				err,
			)
			retries++
		} else {
			log.Log.WithFields(logrus.Fields{
				"machine": info.machine,
			}).Info("machine offline in chain contract success with hash ", hash)

			do.SendOnlineNotify(info.machine, false)
			break
		}
	}

	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel2()
	db.MDB.OfflineMachine(ctx2, info.machine, time.Now())
}

func (do *delayOffline) SendOnlineNotify(machine types.MachineKey, isOnline bool) {
	onr := types.OfflineNotifyRequest{
		MachineId: machine.MachineId,
		IsOnline:  isOnline,
	}
	jsonData, err := json.Marshal(onr)
	if err != nil {
		log.Log.WithFields(logrus.Fields{
			"machine": machine,
			"online":  isOnline,
		}).Errorf("failed to send online notify request: %v", err)
	}

	const maxRetries = 3
	retries := 0
	for retries < maxRetries && machine.Project == "DeepLinkEVM" {
		resp, err := http.Post(
			do.notifyApi,
			"application/json",
			bytes.NewBuffer(jsonData),
		)
		if err != nil || resp.StatusCode != 200 {
			log.Log.WithFields(logrus.Fields{
				"machine": machine,
				"online":  isOnline,
			}).Errorf("failed to send online notify request %v times: %v", retries, err)
		} else {
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Log.WithFields(logrus.Fields{
					"machine": machine,
					"online":  isOnline,
				}).Errorf("failed to send online notify request %v times: %v", retries, err)
			} else {
				result := types.OfflineNotifyResponse{}
				if err := json.Unmarshal(body, &result); err != nil {
					log.Log.WithFields(logrus.Fields{
						"machine": machine,
						"online":  isOnline,
					}).Errorf("failed to send online notify request %v times: %v", retries, err)
				} else {
					if result.Code == 1 {
						log.Log.WithFields(logrus.Fields{
							"machine": machine,
							"online":  isOnline,
						}).Infof("send online notify request success %v %v", result.Success, result.Msg)
						break
					} else {
						log.Log.WithFields(logrus.Fields{
							"machine": machine,
							"online":  isOnline,
						}).Errorf("failed to send online notify request %v times: %v %v", retries, result.Code, result.Msg)
					}
				}
			}
		}
		retries++
	}
}

package ws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"

	"AIL2-ProveNode/db"
	"AIL2-ProveNode/ail2"
	"AIL2-ProveNode/ail2/calculator"
	"AIL2-ProveNode/log"
	"AIL2-ProveNode/types"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 30 * time.Second // 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		if r.Method != "GET" {
			return false
		}
		if r.URL.Path != "/echo" && r.URL.Path != "/websocket" {
			return false
		}
		return true
	},
} // use default options

type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send     chan envelope
	sendDone chan bool

	MachineKey  types.MachineKey
	StakingType types.StakingType
	ClientIP    string
	ClientID    string
}

type envelope struct {
	t   int
	msg []byte
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump(ctx context.Context) {
	log.Log.WithFields(logrus.Fields{
		"uuid":    c.ClientID,
		"machine": c.MachineKey,
	}).Info("read goroutine started")
	// c.hub.wg.Add(1)
	defer func() {
		// c.hub.wg.Done()
		c.hub.wsConns.Delete(c)
		close(c.sendDone)
		close(c.send)
		c.conn.Close()

		log.Log.WithFields(logrus.Fields{
			"uuid":    c.ClientID,
			"machine": c.MachineKey,
		}).Info("read goroutine stopped")
	}()
	// c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	// c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	c.conn.SetPingHandler(func(appData string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		log.Log.WithFields(logrus.Fields{
			"uuid":    c.ClientID,
			"machine": c.MachineKey,
		}).Debug("ping handler")
		err := c.conn.WriteControl(websocket.PongMessage, []byte(appData), time.Now().Add(writeWait))
		if err == websocket.ErrCloseSent {
			return nil
		} else if e, ok := err.(net.Error); ok && e.Temporary() {
			return nil
		}
		return err
	})
	for {
		mt, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Log.WithFields(logrus.Fields{
					"uuid":    c.ClientID,
					"machine": c.MachineKey,
				}).Error("read error: ", err)
			}
			break
		}
		log.Log.WithFields(logrus.Fields{
			"uuid":    c.ClientID,
			"machine": c.MachineKey,
		}).Infof("recv message: %v %s", mt, message)

		// message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		// c.hub.broadcast <- message
		req := &types.WsRequest{}
		if err := json.Unmarshal(message, req); err != nil {
			log.Log.WithFields(logrus.Fields{
				"uuid":    c.ClientID,
				"machine": c.MachineKey,
			}).Error("parse request failed: ", err)
			c.WriteResponse(&types.WsResponse{
				WsHeader: types.WsHeader{
					Version:   0,
					Timestamp: time.Now().Unix(),
					Id:        0,
					Type:      0,
					PubKey:    []byte(""),
					Sign:      []byte(""),
				},
				Code:    uint32(types.ErrCodeParam),
				Message: "parse request failed",
				Body:    []byte(""),
			})
			continue
		}

		// go c.handleRequest(ctx, req)
		c.handleRequest(ctx, req)
	}

	if c.MachineKey.MachineId != "" {
		db.MDB.MachineDisconnected(ctx, c.MachineKey)
		mi, err := db.MDB.GetMachineInfo(ctx, c.MachineKey)
		if err == nil && mi.CalcPoint != 0 && !c.hub.closed() {
			c.hub.do.diconnect <- delayOfflineChanInfo{
				machine:        c.MachineKey,
				disconnectTime: time.Now(),
				stakingType:    c.StakingType,
			}
		}
		// pm.DeleteMetrics(machine)
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump(ctx context.Context) {
	log.Log.WithFields(logrus.Fields{
		"uuid":    c.ClientID,
		"machine": c.MachineKey,
	}).Info("write goroutine started")
	c.hub.wg.Add(1)
	defer func() {
		c.hub.wg.Done()
		c.conn.Close()

		log.Log.WithFields(logrus.Fields{
			"uuid":    c.ClientID,
			"machine": c.MachineKey,
		}).Info("write goroutine stopped")
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if message.t == websocket.CloseMessage {
				c.conn.WriteMessage(websocket.CloseMessage, message.msg)
				return
			}

			w, err := c.conn.NextWriter(message.t) //websocket.TextMessage
			if err != nil {
				return
			}
			w.Write(message.msg)

			// Add queued chat messages to the current websocket message.
			// n := len(c.send)
			// for i := 0; i < n; i++ {
			// 	w.Write(newline)
			// 	w.Write(<-c.send)
			// }

			if err := w.Close(); err != nil {
				return
			}
		case <-ctx.Done():
			c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			return
		}
	}
}

func (c *Client) WriteResponse(res *types.WsResponse) error {
	resBytes, err := json.Marshal(res)
	if err != nil {
		log.Log.WithFields(logrus.Fields{
			"uuid":    c.ClientID,
			"machine": c.MachineKey,
		}).Errorf("marshal reponse with id %v failed: %v", res.Id, err)
		return err
	}

	c.WriteEnvelope(envelope{t: websocket.TextMessage, msg: resBytes})

	return nil
}

func (c *Client) WriteEnvelope(message envelope) {
	select {
	case <-c.sendDone:
		log.Log.WithFields(logrus.Fields{
			"uuid":    c.ClientID,
			"machine": c.MachineKey,
		}).Info("Sender received done signal, exiting")
	default:
		c.send <- message
	}
}

func (c *Client) RespondRequest(req *types.WsRequest, code uint32, message string, body []byte) error {
	return c.WriteResponse(&types.WsResponse{
		WsHeader: types.WsHeader{
			Version:   0,
			Timestamp: time.Now().Unix(),
			Id:        req.Id,
			Type:      req.Type,
			PubKey:    []byte(""),
			Sign:      []byte(""),
		},
		Code:    code,
		Message: message,
		Body:    body,
	})
}

func (c *Client) handleRequest(ctx context.Context, req *types.WsRequest) {
	var (
		code    uint32
		message string
		body    []byte
	)
	switch req.Type {
	case uint32(types.WsMtOnline):
		code, message, body = c.handleOnlineRequest(ctx, req)
	case uint32(types.WsMtDeepLinkMachineInfoST):
		code, message, body = c.handleDeepLinkMachineInfoSTRequest(ctx, req)
	case uint32(types.WsMtDeepLinkMachineInfoBW):
		code, message, body = c.handleDeepLinkMachineInfoBandwidthRequest(ctx, req)
	default:
		code = uint32(types.ErrCodeParam)
		message = "unknowned request message type"
		body = []byte("")
	}
	if code == 0 {
		log.Log.WithFields(logrus.Fields{
			"uuid":    c.ClientID,
			"machine": c.MachineKey,
		}).Info(message)
	} else {
		log.Log.WithFields(logrus.Fields{
			"uuid":    c.ClientID,
			"machine": c.MachineKey,
		}).Error(message)
	}
	c.RespondRequest(req, code, message, body)
}

func (c *Client) handleOnlineRequest(ctx context.Context, req *types.WsRequest) (uint32, string, []byte) {
	if c.MachineKey.MachineId != "" {
		return uint32(types.ErrCodeOnline), "machine has been online, repeated requests", []byte("")
	}

	onlineReq := &types.WsOnlineRequest{}
	if err := json.Unmarshal(req.Body, onlineReq); err != nil {
		return uint32(types.ErrCodeParam), fmt.Sprintf("parse online request failed: %f", err), []byte("")
	}

	if onlineReq.MachineId == "" {
		return uint32(types.ErrCodeParam), "machine id is empty", []byte("")
	}

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	if db.MDB.IsMachineConnected(ctx, onlineReq.MachineKey) {
		return uint32(types.ErrCodeOnline), "machine has been online, repeated connection", []byte("")
	}

	isOnline, isRegistered, err := dbc.DbcChain.GetMachineState(
		ctx,
		onlineReq.Project,
		onlineReq.MachineId,
		onlineReq.StakingType,
	)
	if err != nil {
		return uint32(types.ErrCodeOnline), fmt.Sprintf("failed to get machine state %v", err), []byte("")
	}
	if !isRegistered {
		return uint32(types.ErrCodeOnline), "machine not registered", []byte("")
	}

	if _, err := db.MDB.GetMachineInfo(ctx, onlineReq.MachineKey); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			if err := db.MDB.RegisterMachine(ctx, onlineReq.MachineKey, onlineReq.StakingType); err != nil {
				return uint32(types.ErrCodeOnline), "failed to register machine in mongo", []byte("")
			}
		} else {
			return uint32(types.ErrCodeOnline), "failed to get registered machine info", []byte("")
		}
	}

	if !isOnline {
		if hash, err := dbc.DbcChain.Report(
			ctx,
			types.MachineOnline,
			onlineReq.StakingType,
			onlineReq.Project,
			onlineReq.MachineId,
		); err != nil {
			errMsg := fmt.Sprintf(
				"machine online in chain contract failed with hash %v because of %v",
				hash,
				err,
			)
			return uint32(types.ErrCodeDbcChain), errMsg, []byte("")
		} else {
			log.Log.WithFields(logrus.Fields{
				"uuid":    c.ClientID,
				"machine": onlineReq,
			}).Info("machine online in chain contract success with hash ", hash)
			go c.hub.do.SendOnlineNotify(onlineReq.MachineKey, true)
		}
	}

	if err := db.MDB.MachineConnected(ctx, onlineReq.MachineKey); err != nil {
		return uint32(types.ErrCodeDatabase), fmt.Sprintf("insert online database failed: %v", err), []byte("")
	}

	c.MachineKey = onlineReq.MachineKey
	c.StakingType = onlineReq.StakingType
	// c.hub.wsConns.Store(c, struct{}{})
	c.hub.do.connect <- delayOfflineChanInfo{
		machine:        c.MachineKey,
		disconnectTime: time.Now(),
		stakingType:    c.StakingType,
	}
	return 0, "machine online success", []byte("")
}

func (c *Client) handleDeepLinkMachineInfoSTRequest(ctx context.Context, req *types.WsRequest) (uint32, string, []byte) {
	if c.MachineKey.MachineId == "" {
		return uint32(types.ErrCodeMachineInfo), "machine id is empty, need send online request first", []byte("")
	}

	miReq := types.DeepLinkMachineInfoST{}
	if err := json.Unmarshal(req.Body, &miReq); err != nil {
		return uint32(types.ErrCodeParam), fmt.Sprintf("parse machine info request failed: %v", err), []byte("")
	}

	if err := miReq.Validate(); err != nil {
		return uint32(types.ErrCodeParam), err.Error(), []byte("")
	}
	miReq.ClientIP = c.ClientIP

	ctx, cancel := context.WithTimeout(ctx, 80*time.Second)
	defer cancel()
	mi, err := db.MDB.GetMachineInfo(ctx, c.MachineKey)
	if err != nil {
		return uint32(types.ErrCodeDatabase), fmt.Sprintf("get machine info from database failed: %v", err), []byte("")
	} else if mi.MDBDeepLinkMachineInfoST.CalcPoint == 0 {
		calcPoint, err := calculator.CalculatePointExactFromReport(
			miReq.GPUNames,
			miReq.GPUMemoryTotal,
			miReq.MemoryTotal,
		)
		if calcPoint == 0 {
			log.Log.WithFields(logrus.Fields{
				"uuid":    c.ClientID,
				"machine": c.MachineKey,
			}).Errorf("calculate gpu point from report %v failed %v", miReq, err)
			calcPoint, err = calculator.CalculatePointFuzzyFromReport(
				miReq.GPUNames,
				miReq.MemoryTotal,
			)
		}

		if calcPoint == 0 {
			log.Log.WithFields(logrus.Fields{
				"uuid":    c.ClientID,
				"machine": c.MachineKey,
			}).Errorf("calculate gpu point from report %v failed %v", miReq, err)
			return uint32(types.ErrCodeMachineInfo), fmt.Sprintf("calculate gpu point failed %v", err), []byte("")
		} else {
			log.Log.WithFields(logrus.Fields{
				"uuid":    c.ClientID,
				"machine": c.MachineKey,
			}).Infof("calculate gpu point from reported machine info %v => %v", miReq, calcPoint)

			loc, err := db.GetPositionOfIP(c.ClientIP)
			if err != nil {
				log.Log.WithFields(logrus.Fields{
					"uuid":    c.ClientID,
					"machine": c.MachineKey,
				}).Errorf("get location err %v from ip address %v", err, c.ClientIP)
			} else {
				log.Log.WithFields(logrus.Fields{
					"uuid":    c.ClientID,
					"machine": c.MachineKey,
				}).Infof("get location (%f, %f) from ip address %v", loc.Longitude, loc.Latitude, c.ClientIP)
			}

			if hash, err := dbc.DbcChain.SetDeepLinkMachineInfoST(
				ctx,
				c.MachineKey,
				miReq,
				int64(calcPoint*10000),
				loc.Longitude,
				loc.Latitude,
				loc.Region,
			); err != nil {
				errMsg := fmt.Sprintf(
					"set machine info in chain contract with hash %v failed: %v",
					hash,
					err,
				)
				return uint32(types.ErrCodeDbcChain), errMsg, []byte("")
			} else {
				log.Log.WithFields(logrus.Fields{
					"uuid":    c.ClientID,
					"machine": c.MachineKey,
				}).Info("set machine info in chain contract success with hash ", hash)

				if err := db.MDB.SetMachineInfo(
					ctx,
					c.MachineKey,
					&types.MDBDeepLinkMachineInfoST{
						DeepLinkMachineInfoST: miReq,
						CalcPoint:             calcPoint,
						Longitude:             loc.Longitude,
						Latitude:              loc.Latitude,
						Region:                loc.Region,
					},
					nil,
				); err != nil {
					return uint32(types.ErrCodeDbcChain), fmt.Sprintf("set machine info in database failed %v", err), []byte("")
				} else {
					log.Log.WithFields(logrus.Fields{
						"uuid":    c.ClientID,
						"machine": c.MachineKey,
					}).Info("set machine info in database success")
				}
			}
		}
	} else {
		log.Log.WithFields(logrus.Fields{
			"uuid":    c.ClientID,
			"machine": c.MachineKey,
		}).WithField("machine info", miReq).Info("get machine info success and calcpoint ", mi.MDBDeepLinkMachineInfoST.CalcPoint)
	}

	if err := db.MDB.AddMachineTM(
		ctx,
		types.MDBMachineTM{
			Timestamp:                    time.UnixMilli(req.Timestamp),
			Machine:                      c.MachineKey,
			DeepLinkMachineInfoST:        miReq,
			DeepLinkMachineInfoBandwidth: types.DeepLinkMachineInfoBandwidth{},
		},
	); err != nil {
		return uint32(types.ErrCodeDatabase), fmt.Sprintf("add machine tm in database failed %v", err), []byte("")
	}

	// pm.SetMetrics(nodeId, miReq)
	return 0, "add machine tm success", []byte("")
}

func (c *Client) handleDeepLinkMachineInfoBandwidthRequest(ctx context.Context, req *types.WsRequest) (uint32, string, []byte) {
	if c.MachineKey.MachineId == "" {
		return uint32(types.ErrCodeMachineInfo), "machine id is empty, need send online request first", []byte("")
	}

	miReq := types.DeepLinkMachineInfoBandwidth{}
	if err := json.Unmarshal(req.Body, &miReq); err != nil {
		return uint32(types.ErrCodeParam), fmt.Sprintf("parse machine info bandwidth request failed: %v", err), []byte("")
	}

	if err := miReq.Validate(); err != nil {
		return uint32(types.ErrCodeParam), err.Error(), []byte("")
	}
	miReq.ClientIP = c.ClientIP

	ctx, cancel := context.WithTimeout(ctx, 80*time.Second)
	defer cancel()
	mi, err := db.MDB.GetMachineInfo(ctx, c.MachineKey)
	if err != nil {
		return uint32(types.ErrCodeDatabase), fmt.Sprintf("get machine info bandwidth from database failed: %v", err), []byte("")
	} else if mi.MDBDeepLinkMachineInfoBandwidth.Bandwidth == 0 {
		loc, err := db.GetPositionOfIP(c.ClientIP)
		if err != nil {
			log.Log.WithFields(logrus.Fields{
				"uuid":    c.ClientID,
				"machine": c.MachineKey,
			}).Errorf("get location err %v from ip address %v", err, c.ClientIP)
		} else {
			log.Log.WithFields(logrus.Fields{
				"uuid":    c.ClientID,
				"machine": c.MachineKey,
			}).Infof("get location (%f, %f) from ip address %v", loc.Longitude, loc.Latitude, c.ClientIP)
		}
		region := db.GetBandwidthRegion(&loc)

		if hash, err := dbc.DbcChain.SetDeepLinkMachineInfoBandwidth(
			ctx,
			c.MachineKey,
			miReq,
			region,
		); err != nil {
			errMsg := fmt.Sprintf(
				"set machine info bandwidth in chain contract with hash %v failed: %v",
				hash,
				err,
			)
			return uint32(types.ErrCodeDbcChain), errMsg, []byte("")
		} else {
			log.Log.WithFields(logrus.Fields{
				"uuid":    c.ClientID,
				"machine": c.MachineKey,
			}).Info("set machine info bandwidth in chain contract success with hash ", hash)

			if err := db.MDB.SetMachineInfo(
				ctx,
				c.MachineKey,
				nil,
				&types.MDBDeepLinkMachineInfoBandwidth{
					DeepLinkMachineInfoBandwidth: miReq,
					Region:                       region,
				},
			); err != nil {
				return uint32(types.ErrCodeDbcChain), fmt.Sprintf("set machine info bandwidth in database failed %v", err), []byte("")
			} else {
				log.Log.WithFields(logrus.Fields{
					"uuid":    c.ClientID,
					"machine": c.MachineKey,
				}).Info("set machine info bandwidth in database success")
			}
		}
	} else {
		log.Log.WithFields(logrus.Fields{
			"uuid":    c.ClientID,
			"machine": c.MachineKey,
		}).WithField("machine info", miReq).Info(
			"get machine info bandwidth success and bandwidth ",
			mi.MDBDeepLinkMachineInfoBandwidth.Bandwidth,
		)
	}

	if err := db.MDB.AddMachineTM(
		ctx,
		types.MDBMachineTM{
			Timestamp:                    time.UnixMilli(req.Timestamp),
			Machine:                      c.MachineKey,
			DeepLinkMachineInfoST:        types.DeepLinkMachineInfoST{},
			DeepLinkMachineInfoBandwidth: miReq,
		},
	); err != nil {
		return uint32(types.ErrCodeDatabase), fmt.Sprintf("add machine bandwidth tm in database failed %v", err), []byte("")
	}

	// pm.SetMetrics(nodeId, miReq)
	return 0, "add machine bandwidth tm success", []byte("")
}

func Ws2(hub *Hub, ctx *gin.Context, wsCtx context.Context) {
	w, r := ctx.Writer, ctx.Request
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Upgrade to websocket failed", http.StatusUpgradeRequired)
		log.Log.Error("Upgrade to websocket failed: ", err)
		return
	}

	randomUUID, err := uuid.NewRandom()
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			types.BaseHttpResponse{
				Code:    int(types.ErrCodeUUID),
				Message: types.ErrCodeUUID.String(),
			},
		)
		log.Log.Error("generate uuid for websocket connection failed: ", err)
		return
	}

	client := &Client{
		hub:      hub,
		conn:     c,
		send:     make(chan envelope, 32),
		sendDone: make(chan bool),
		ClientIP: ctx.ClientIP(),
		ClientID: randomUUID.String(),
	}
	hub.wsConns.Store(
		client,
		struct{}{},
	)

	go client.writePump(wsCtx)
	// go client.readPump(context.Background())
	client.readPump(context.Background())
}

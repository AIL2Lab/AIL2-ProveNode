package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/showwin/speedtest-go/speedtest"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"

	"AIL2-ProveNode/types"
)

var addr = flag.String("addr", "health0.deepbrainchain.org", "websocket service address")
var wallet = flag.String("wallet", "", "EVM wallet address (42 characters)")

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 10 * time.Second // 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512

	configFile = "config.json"
)

type envelope struct {
	t   int
	msg []byte
}

type systemInfo struct {
	CpuCores    int32  `json:"cpu_cores" bson:"cpu_cores,omitempty"`
	MemoryTotal int64  `json:"memory_total" bson:"memory_total,omitempty"` // GB
	Hdd         int64  `json:"hdd" bson:"hdd,omitempty"`
	Bandwidth   int32  `json:"bandwidth" bson:"bandwidth,omitempty"`
	Wallet      string `json:"wallet" bson:"wallet,omitempty"`
}

type Config struct {
	MachineId  string `json:"machine_id"`
	PrivateKey string `json:"private_key"`
	Wallet     string `json:"wallet"`
}

func GenMachineId() (string, string, error) {
	// use secp256k1 gen private key
	privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate key pair: %v", err)
	}

	privateKeyHex := "0x" + hex.EncodeToString(crypto.FromECDSA(privateKey))

	publicKey := privateKey.PublicKey
	publicKeyBytes := crypto.FromECDSAPub(&publicKey)
	if len(publicKeyBytes) > 32 {
		publicKeyBytes = publicKeyBytes[1:33]
	}

	publicKeyHex := hex.EncodeToString(publicKeyBytes)

	return publicKeyHex, privateKeyHex, nil
}

func GetSystemStats() systemInfo {
	// get disk size
	diskStat, _ := disk.Usage("/")
	diskSizeGB := diskStat.Total / 1024 / 1024 / 1024

	// get CPU cores
	cpuCores, _ := cpu.Counts(true)

	// get mem size
	memStat, _ := mem.VirtualMemory()
	memSizeGB := memStat.Total / 1000 / 1000 / 1000

	// get ISP and IP
	//user, _ := speedtest.FetchUserInfo()

	// upload test
	speedTestClient := speedtest.New()
	serverList, err := speedTestClient.FetchServers()
	if err != nil {
		log.Println("get speedtest server list failed:", err)
	}
	targets, err := serverList.FindServer([]int{})
	if err != nil {
		log.Println("find speedtest server failed:", err)
	}

	var uploadSpeed float64
	for _, server := range targets {
		server.UploadTest()
		uploadSpeed = float64((server.ULSpeed * 8) / 1e6) // B/s 转 Mbps
		break
	}

	return systemInfo{
		Hdd:         int64(diskSizeGB),
		CpuCores:    int32(cpuCores),
		MemoryTotal: int64(memSizeGB),
		Bandwidth:   int32(uploadSpeed),
	}
}

func loadConfig() (*Config, error) {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		log.Println("Config file does not exist, generating new config...")
		publicKeyHex, privateKeyHex, err := GenMachineId()
		if err != nil {
			return nil, err
		}

		// 生成新配置时，不包含 system_info
		config := &Config{
			MachineId:  publicKeyHex,
			PrivateKey: privateKeyHex,
			Wallet:     "", // 默认钱包为空
		}

		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			return nil, err
		}

		if err := os.WriteFile(configFile, configBytes, 0644); err != nil {
			return nil, err
		}

		return config, nil
	}

	configBytes, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(configBytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	if *wallet != "" {
		// 检查钱包地址
		if len(*wallet) != 42 {
			log.Fatal("Error: EVM wallet address must be 42 characters long, including '0x' prefix")
		}

		// 加载或生成配置
		config, err := loadConfig()
		if err != nil {
			log.Fatalf("Failed to load or create config: %v", err)
		}

		// 更新钱包地址
		config.Wallet = *wallet

		// 确保配置文件中不包含 system_info
		cleanConfig := &Config{
			MachineId:  config.MachineId,
			PrivateKey: config.PrivateKey,
			Wallet:     config.Wallet,
		}

		// 保存配置
		configBytes, err := json.MarshalIndent(cleanConfig, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling updated config: %v", err)
		}
		if err := os.WriteFile(configFile, configBytes, 0644); err != nil {
			log.Fatalf("Error writing updated config file: %v", err)
		}

		fmt.Println("Wallet address has been set and saved to config.json")
		fmt.Println("Your machine id: ", config.MachineId)
		os.Exit(0)
	}

	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load or create config: %v", err)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	dialer := websocket.Dialer{
		NetDial: func(network, addr string) (net.Conn, error) {
			return net.Dial("tcp4", addr)
		},
		HandshakeTimeout: 30 * time.Second,
	}

	u := url.URL{Scheme: "wss", Host: *addr, Path: "/websocket"}
	log.Printf("connecting to %s", u.String())

	c, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial error:", err)
	}
	defer c.Close()

	done := make(chan struct{})
	writeQueue := make(chan envelope)
	var reqId uint64 = 0

	// send machine info using go sendOnline(time.Second)
	sendOnline := func(delay time.Duration) {
		time.Sleep(delay)
		onlineReq := &types.WsOnlineRequest{
			MachineKey: types.MachineKey{
				MachineId:   config.MachineId,
				Project:     "DeepLink BandWidth",
				ContainerId: "",
			},
			StakingType: types.Free,
		}
		reqBody, err := json.Marshal(onlineReq)
		if err != nil {
			log.Fatalf("marshal online request body failed: %v", err)
		}
		req := &types.WsRequest{
			WsHeader: types.WsHeader{
				Version:   0,
				Timestamp: time.Now().UnixMilli(),
				Id:        reqId,
				Type:      uint32(types.WsMtOnline),
				PubKey:    []byte(""),
				Sign:      []byte(""),
			},
			Body: reqBody,
		}
		reqBytes, err := json.Marshal(req)
		if err != nil {
			log.Fatalf("marshal online request failed: %v", err)
		}
		select {
		case <-done:
			log.Print("connection already closed")
		default:
			writeQueue <- envelope{t: websocket.TextMessage, msg: reqBytes}
			reqId++
		}
	}

	// send machine info using go sendMachineInfo(time.Second)
	sendMachineInfo := func(delay time.Duration) {
		time.Sleep(delay)
		systemInfo := GetSystemStats()
		machineInfo := &types.DeepLinkMachineInfoBandwidth{
			CpuCores:    systemInfo.CpuCores,
			MemoryTotal: systemInfo.MemoryTotal,
			Hdd:         systemInfo.Hdd,
			Bandwidth:   systemInfo.Bandwidth,
			Wallet:      config.Wallet,
		}
		reqBody, err := json.Marshal(machineInfo)
		if err != nil {
			log.Fatalf("marshal machine info request body failed: %v", err)
		}
		req2 := &types.WsRequest{
			WsHeader: types.WsHeader{
				Version:   0,
				Timestamp: time.Now().UnixMilli(),
				Id:        reqId,
				Type:      uint32(types.WsMtDeepLinkMachineInfoBW),
				PubKey:    []byte(""),
				Sign:      []byte(""),
			},
			Body: reqBody,
		}
		reqBytes, err := json.Marshal(req2)
		if err != nil {
			log.Fatalf("marshal machine info request failed: %v", err)
		}
		select {
		case <-done:
			log.Print("connection already closed")
		default:
			writeQueue <- envelope{t: websocket.TextMessage, msg: reqBytes}
			reqId++
		}
	}
	c.SetPongHandler(func(string) error { c.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	// read websocket connection
	go func() {
		defer close(done)
		for {
			c.SetReadDeadline(time.Now().Add(pongWait))
			_, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("read error: %v", err)
				} else {
					log.Printf("connection closed: %v", err)
				}
				return
			}
			log.Printf("recv: %s", message)

			response := &types.WsResponse{}
			if err := json.Unmarshal(message, response); err != nil {
				log.Printf("parse response failed: %v", err)
				continue
			}
			switch response.Type {
			case uint32(types.WsMtOnline):
				if response.Code == 0 {
					go sendMachineInfo(3 * time.Second)
				} else {
					log.Printf("online failed %v %v", response.Code, response.Message)
					go sendOnline(10 * time.Second)
				}
			case uint32(types.WsMtDeepLinkMachineInfoBW):
				if response.Code == 0 {
					log.Printf("send machine info bandwidth success")
				} else {
					log.Printf("online failed %v %v", response.Code, response.Message)
					go sendMachineInfo(10 * time.Second)
				}
			case uint32(types.WsMtNotify):
				notifyMessage := &types.WsNotifyMessage{}
				if err := json.Unmarshal(response.Body, notifyMessage); err != nil {
					log.Printf("parse notify response failed: %v", err)
				} else {
					if notifyMessage.Unregister.Message != "" {
						log.Printf("machine was unregistered by server")
						return
					}
				}
			}
		}
	}()

	go sendOnline(3 * time.Second)
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			if err := c.WriteControl(websocket.PingMessage, []byte(t.String()), time.Now().Add(writeWait)); err != nil {
				log.Printf("ping websocket failed: %v", err)
				return
			}
		case message, ok := <-writeQueue:
			c.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// sen close message
				c.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if message.t == websocket.CloseMessage {
				c.WriteMessage(websocket.CloseMessage, message.msg)
				return
			}
			c.WriteMessage(message.t, message.msg)
		case <-interrupt:
			log.Println("interrupt")

			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

package ws

import (
	"encoding/json"
	"testing"
	"time"

	"AIL2-ProveNode/types"

	"github.com/gorilla/websocket"
)

// go test -v -timeout 30s -count=1 -run TestWsMachine AIL2-ProveNode/ws
func TestWsMachine(t *testing.T) {
	url := "ws://localhost:9521/websocket"
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("dial %s: %v", url, err)
	}
	defer c.Close()

	if err := c.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(9*time.Second)); err != nil {
		t.Fatalf("ping websocket failed: %v", err)
	}

	done := make(chan struct{})
	resChan := make(chan types.WsResponse)
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				t.Log("read:", err)
				return
			}
			t.Logf("recv: %s", message)

			response := types.WsResponse{}
			if err := json.Unmarshal(message, &response); err != nil {
				println("parse websocket response failed:", err)
				break
			}
			resChan <- response
		}
	}()

	var reqId uint64 = 0

	onlineReq := &types.WsOnlineRequest{
		MachineKey: types.MachineKey{
			MachineId:   "123456789",
			Project:     "deeplink",
			ContainerId: "",
		},
		StakingType: types.ShortTerm,
	}
	reqBody, err := json.Marshal(onlineReq)
	if err != nil {
		t.Fatalf("marshal online request body failed: %v", err)
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
		t.Fatalf("marshal online request failed: %v", err)
	}
	if err := c.WriteMessage(websocket.TextMessage, reqBytes); err != nil {
		t.Fatalf("send websocket message failed: %v", err)
	}
	t.Logf("send: %v", string(reqBytes))

	onlineRes := <-resChan
	if onlineRes.Type != uint32(types.WsMtOnline) || onlineRes.Code != 0 {
		t.Fatal("received online request", onlineRes.Message)
	}

	time.Sleep(2 * time.Second)
	reqId++

	machineInfo := &types.DeepLinkMachineInfoST{
		CpuType:        "13th Gen Intel(R) Core(TM) i5-13400F",
		CpuRate:        2500,
		GPUNames:       []string{"VIDIA GeForce RTX 4060"},
		GPUMemoryTotal: []int32{8},
		// UtilizationGPU: 30,
		MemoryTotal: 24564,
		// MemoryUsed:     22128,
		Wallet: "0x109dd702f055f6752afffb3247ba18d843fab713",
	}
	// machineInfo.Models = append(machineInfo.Models, types.ModelInfo{
	// 	Model: "Codestral-22B-v0.1",
	// })
	reqBody, err = json.Marshal(machineInfo)
	if err != nil {
		t.Fatalf("marshal machine info request body failed: %v", err)
	}
	req2 := &types.WsRequest{
		WsHeader: types.WsHeader{
			Version:   0,
			Timestamp: time.Now().UnixMilli(),
			Id:        reqId,
			Type:      uint32(types.WsMtDeepLinkMachineInfoST),
			PubKey:    []byte(""),
			Sign:      []byte(""),
		},
		Body: reqBody,
	}
	reqBytes, err = json.Marshal(req2)
	if err != nil {
		t.Fatalf("marshal machine info request failed: %v", err)
	}
	if err := c.WriteMessage(websocket.TextMessage, reqBytes); err != nil {
		t.Fatalf("send websocket message failed: %v", err)
	}

	miRes := <-resChan
	if miRes.Type == uint32(types.WsMtDeepLinkMachineInfoST) {
		t.Log("received machine info response", miRes.Code, miRes.Message)
	}

	time.Sleep(2 * time.Second)

	err = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		t.Log("write close:", err)
		return
	}
	select {
	case <-done:
	case <-time.After(time.Second):
	}
}

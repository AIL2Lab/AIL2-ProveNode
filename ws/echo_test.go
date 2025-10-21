package ws

import (
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

// go test -v -timeout 30s -count=1 -run TestWsEcho AIL2-ProveNode/ws
func TestWsEcho(t *testing.T) {
	url := "ws://localhost:9521/echo"
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("dial %s: %v", url, err)
	}
	defer c.Close()

	c.SetPingHandler(func(appData string) error {
		t.Log("ping handler")
		return nil
	})
	c.SetPongHandler(func(appData string) error {
		t.Log("pong handler")
		return nil
	})

	done := make(chan struct{})

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				t.Log("read:", err)
				return
			}
			t.Logf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	pingTicker := time.NewTicker(3 * time.Second)
	defer pingTicker.Stop()

	for {
		select {
		case <-done:
			t.Log("done")
			return
		case tm := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(tm.String()))
			if err != nil {
				t.Log("write:", err)
				return
			}
		case <-pingTicker.C:
			// if err := ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait)); err != nil {
			c.SetWriteDeadline(time.Now().Add(9 * time.Second))
			if err := c.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				t.Log("ping:", err)
			}
		case <-interrupt:
			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				t.Log("write close:", err)
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

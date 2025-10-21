package ws

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func Echo(ctx *gin.Context) {
	w, r := ctx.Writer, ctx.Request
	log.Print("IsWebSocketUpgrade:", websocket.IsWebSocketUpgrade(r))
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	// c.SetPingHandler(func(appData string) error {
	// 	log.Println("ping handler")
	// 	return nil
	// })
	c.SetPongHandler(func(appData string) error {
		log.Println("pong handler")
		return nil
	})

	c.SetReadDeadline(time.Now().Add(10 * time.Second))
	c.SetPingHandler(func(appData string) error {
		c.SetReadDeadline(time.Now().Add(10 * time.Second))
		log.Println("ping handler")
		return nil
	})

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv message: {type: %v, content: %s}", mt, string(message))
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

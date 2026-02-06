package websocket

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}

func ServeWS(hub *Hub) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {

		client := &Client{
			Conn: c,
			Send: make(chan []byte),
		}
		hub.Register <- client

		go func() {
			for msg := range client.Send {
				c.WriteMessage(websocket.TextMessage, msg)
			}
		}()
		for {
			if _, _, err := c.ReadMessage(); err != nil {
				break
			}
		}
		hub.Unregister <- client
	})
}

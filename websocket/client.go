package websocket

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Client struct {
	Conn      *websocket.Conn
	Send      chan []byte
	VehicleID uint
}

func ServeWS(hub *Hub) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		vehicleIDStr := c.Query("vehicle_id")
		vid, _ := strconv.ParseUint(vehicleIDStr, 10, 64)

		client := &Client{
			Conn:      c,
			Send:      make(chan []byte),
			VehicleID: uint(vid),
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

package websocket

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
)

type Client struct {
	Conn      *websocket.Conn
	Send      chan []byte
	UserID    uint
	VehicleID uint
}

func ServeWS(hub *Hub) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {

		token := c.Query("token")
		claims, err := helpers.ValidateJWT(token)
		if err != nil {
			return
		}
		userID := claims.UserID
		vehicleIDStr := c.Query("vehicle_id")
		vid, _ := strconv.ParseUint(vehicleIDStr, 10, 64)

		client := &Client{
			Conn:      c,
			Send:      make(chan []byte),
			UserID:    userID,
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

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/websocket"
)

func WebsocketRoutes(app *fiber.App, hub *websocket.Hub) {
	app.Get("/ws", websocket.ServeWS(hub))
}

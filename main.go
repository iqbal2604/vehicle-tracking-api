package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Ping")

	})

	err := app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}

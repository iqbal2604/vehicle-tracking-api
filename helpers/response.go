package helpers

import "github.com/gofiber/fiber/v2"

func ErrorResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"error": message,
	})
}

func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.JSON(data)
}

func Error(c *fiber.Ctx, data interface{}) error {
	return c.JSON(data)
}

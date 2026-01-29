package helpers

import "github.com/gofiber/fiber/v2"

func ErrorResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"error": message,
	})
}

func SuccessResponse(c *fiber.Ctx, message string) error {
	return c.JSON(fiber.Map{
		"message": message,
	})
}

package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestStream(t *testing.T) {
	app := fiber.New()

	app.Get("api/gps/stream/:vehicle_id", func(c *fiber.Ctx) error {
		c.Set("Content_Type", "text/event-stream")
		return c.SendString("data: {\"lat\":4}\n\n")
	})

	req := httptest.NewRequest(
		http.MethodGet, "/api/gps/stream/4",
		nil,
	)

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bytes, _ := io.ReadAll(resp.Body)
	bodyString := string(bytes)

	assert.Contains(t, bodyString, "data:")
}

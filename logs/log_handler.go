package logs

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
)

type LogHandler struct {
	service LogService
}

func NewLogHandler(service LogService) *LogHandler {
	return &LogHandler{service: service}
}

func (h *LogHandler) GetRecentLogs(c *fiber.Ctx) error {
	limitStr := c.Query("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		return helpers.ErrorResponse(c, 400, "Invalid Limit Parameter")
	}

	logs, err := h.service.GetRecent(limit)
	if err != nil {
		return helpers.ErrorResponse(c, 500, "Failed to Retrieve Log")
	}

	return helpers.SuccessResponse(c, fiber.Map{
		"logs": logs,
	})
}

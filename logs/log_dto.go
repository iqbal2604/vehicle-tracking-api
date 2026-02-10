package logs

import (
	"time"
)

type LogResponse struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	Action    string    `json:"action"`
	UserID    *uint     `json:"user_id"`
	TargetID  *uint     `json:"target_id"`
	Meta      string    `json:"meta"`
	IPAddress string    `json:"ip_address"`
	CreatedAt time.Time `json:"created-at"`
}

func ToLogResponse(log Log) LogResponse {
	return LogResponse{
		ID:        log.ID,
		Type:      log.Type,
		Action:    log.Action,
		UserID:    log.UserID,
		TargetID:  log.TargetID,
		Meta:      log.Meta,
		IPAddress: log.IPAddress,
		CreatedAt: log.CreatedAt,
	}
}

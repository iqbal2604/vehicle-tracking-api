package logs

import (
	"time"
)

type LogResponse struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	Action    string    `json:"action"`
	UserID    *uint     `json:"user_id"`
	UserName  string    `json:"user_name"`
	TargetID  *uint     `json:"target_id"`
	Meta      string    `json:"meta"`
	IPAddress string    `json:"ip_address"`
	CreatedAt time.Time `json:"created-at"`
}

func ToLogResponse(log Log) LogResponse {

	var userName string

	if log.UserID != nil {
		userName = log.User.Name
	}
	return LogResponse{
		ID:        log.ID,
		Type:      log.Type,
		Action:    log.Action,
		UserID:    log.UserID,
		UserName:  userName,
		TargetID:  log.TargetID,
		Meta:      log.Meta,
		IPAddress: log.IPAddress,
		CreatedAt: log.CreatedAt,
	}
}

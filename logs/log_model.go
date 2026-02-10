package logs

import "time"

type Log struct {
	ID        uint   `gorm:"primaryKey"`
	Type      string `gorm:"size:50"`
	Action    string `gorm:"size:100"`
	UserID    *uint
	TargetID  *uint
	Meta      string `gorm:"type:text"`
	IPAddress string `gorm:"size:45"`
	CreatedAt time.Time
}

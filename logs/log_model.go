package logs

import (
	"time"

	domain "github.com/iqbal2604/vehicle-tracking-api/models/domain"
)

type Log struct {
	ID        uint   `gorm:"primaryKey"`
	Type      string `gorm:"size:50"`
	Action    string `gorm:"size:100"`
	UserID    *uint
	User      domain.User `gorm:foreignKey:UserID`
	TargetID  *uint
	Meta      string `gorm:"type:text"`
	IPAddress string `gorm:"size:45"`
	CreatedAt time.Time
}

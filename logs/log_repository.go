package logs

import "gorm.io/gorm"

type LogRepository interface {
	Create(log *Log) error
	FindAll(limit int) ([]Log, error)
}

type logRepository struct {
	DB *gorm.DB
}

func NewLogRepository(db *gorm.DB) LogRepository {
	return &logRepository{DB: db}
}

func (r *logRepository) Create(log *Log) error {
	return r.DB.Create(log).Error
}

func (r *logRepository) FindAll(limit int) ([]Log, error) {
	var logs []Log
	err := r.DB.Preload("User").Order("created_at desc").Limit(limit).Find(&logs).Error
	return logs, err
}

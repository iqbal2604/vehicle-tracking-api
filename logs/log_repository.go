package logs

import "gorm.io/gorm"

type LogRepository interface {
	Create(log *Log) error
	FindAll(limit, offset int) ([]Log, error)
	CountAll() (int64, error)
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

func (r *logRepository) FindAll(limit, offset int) ([]Log, error) {
	var logs []Log
	err := r.DB.Preload("User").Order("created_at desc").Limit(limit).Offset(offset).Find(&logs).Error
	return logs, err
}

func (r *logRepository) CountAll() (int64, error) {
	var count int64
	err := r.DB.Model(&Log{}).Count(&count).Error
	return count, err
}

package logs

import "time"

type LogServiceImpl struct {
	repo LogRepository
}

func NewLogServiceImpl(repo LogRepository) LogService {
	return &LogServiceImpl{repo: repo}
}

func (s *LogServiceImpl) LogAuth(action string, userID *uint, meta string, ip string) {
	log := &Log{
		Type:      "auth",
		Action:    action,
		UserID:    userID,
		Meta:      meta,
		IPAddress: ip,
		CreatedAt: time.Now(),
	}
	s.repo.Create(log)
}

func (s *LogServiceImpl) LogAdmin(action string, adminID uint, targetID *uint, meta string) {
	log := &Log{
		Type:      "admin",
		Action:    action,
		UserID:    &adminID,
		TargetID:  targetID,
		Meta:      meta,
		IPAddress: "",
		CreatedAt: time.Now(),
	}

	s.repo.Create(log)
}

func (s *LogServiceImpl) LogSystem(action string, meta string) {
	log := &Log{
		Type:      "system",
		Action:    action,
		Meta:      meta,
		IPAddress: "",
		CreatedAt: time.Now(),
	}

	s.repo.Create(log)
}

func (s *LogServiceImpl) GetLogs(page, limit int) ([]LogResponse, int64, error) {
	offset := (page - 1) * limit

	var logs []Log
	var totalCounts int64
	var errLogs, errCount error

	done := make(chan bool, 2)

	go func() {
		logs, errLogs = s.repo.FindAll(limit, offset)
		done <- true
	}()

	go func() {
		totalCounts, errCount = s.repo.CountAll()
		done <- true
	}()

	for i := 0; i < 2; i++ {
		<-done
	}

	if errLogs != nil {
		return nil, 0, errLogs
	}

	if errCount != nil {
		return nil, 0, errCount
	}

	var response []LogResponse
	for _, log := range logs {
		response = append(response, ToLogResponse(log))
	}

	return response, totalCounts, nil

}

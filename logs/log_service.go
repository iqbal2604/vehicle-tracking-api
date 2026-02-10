package logs

type LogService interface {
	LogAuth(action string, userID *uint, meta string, ip string)
	LogAdmin(action string, adminID uint, targetID *uint, meta string)
	LogSystem(action string, meta string)
	GetRecent(limit int) ([]LogResponse, error)
}

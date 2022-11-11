package rfsl

type LogLevel int

const (
	LOG_LVL_ALL LogLevel = iota
	LOG_LVL_TRACE
	LOG_LVL_DEBUG
	LOG_LVL_INFO
	LOG_LVL_WARN
	LOG_LVL_ERROR
	LOG_LVL_PANIC
)

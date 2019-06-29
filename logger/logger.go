package logger

import "github.com/chenzr/gin-web/config"

var l *Logger

func init() {
	cfg := config.GetConfig()

	time := true
	file := true
	debug := cfg.DebugLog
	trace := cfg.TraceLog
	colors := true
	pid := false

	if cfg.LogFile == "" {
		l = NewStdLogger(time, file, debug, trace, colors, pid)
	} else {
		l = NewFileLogger(cfg.LogFile, time, file, debug, trace, pid)
	}
}

func GetLogger() *Logger {
	return l
}

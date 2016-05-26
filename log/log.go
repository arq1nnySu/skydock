package log

import (
	"log"
	"os"
)

const (
	FATAL = iota
	ERROR
	INFO
	DEBUG
)

type (
	Logger interface{}
)

var (
	loggers = map[int]*log.Logger {
		FATAL: log.New(os.Stdout, "[fatal] ", 0),
		ERROR: log.New(os.Stdout, "[error] ", 0),
		INFO:  log.New(os.Stdout, "[info] ",  0),
		DEBUG: log.New(os.Stdout, "[debug] ", 0),
	}
)

func Logf(severity int, format string, opts ...interface{}) {
	loggers[severity].Printf(format, opts)
}

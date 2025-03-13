package word

import "log"

type Logger interface {
	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

type DefaultLogger struct {
	LogLevel int
}

func (l *DefaultLogger) Errorf(format string, args ...interface{}) {
	log.Printf("ERROR: "+format, args...)
}

func (l *DefaultLogger) Infof(format string, args ...interface{}) {
	if l.LogLevel <= LogLevelInfo {
		log.Printf("INFO: "+format, args...)
	}
}

func (l *DefaultLogger) Debugf(format string, args ...interface{}) {
	if l.LogLevel <= LogLevelDebug {
		log.Printf("DEBUG: "+format, args...)
	}
}

func (c *Client) SetLogger(logger Logger) {
	c.logger = logger
}

func (l *DefaultLogger) SetLogLevel(level int) {
	l.LogLevel = level
}

var (
	LogLevelDebug = 1
	LogLevelInfo  = 2
	LogLevelError = 3
)

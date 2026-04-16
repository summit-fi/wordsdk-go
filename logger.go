package word

import "log"

type LoggerLevel string

const (
	LogLevelError LoggerLevel = "error"
	LogLevelInfo  LoggerLevel = "info"
	LogLevelDebug LoggerLevel = "debug"
)

type Logger interface {
	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

type DefaultLogger struct {
	LoggerLevel
}

func (l *DefaultLogger) Errorf(format string, args ...interface{}) {
	log.Printf("ERROR: "+format, args...)
}

func (l *DefaultLogger) Infof(format string, args ...interface{}) {
	if l.LoggerLevel == LogLevelInfo || l.LoggerLevel == LogLevelDebug {
		log.Printf("INFO: "+format, args...)
	}
}

func (l *DefaultLogger) Debugf(format string, args ...interface{}) {
	if l.LoggerLevel == LogLevelDebug {
		log.Printf("DEBUG: "+format, args...)
	}
}

func (c *Client) SetLogger(logger Logger) {
	c.logger = logger
}

func (l *DefaultLogger) SetLogLevel(level LoggerLevel) {
	l.LoggerLevel = level
}

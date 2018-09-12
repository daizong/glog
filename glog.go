package glog

import (
	"io"
	"log"
)

const (
	DEBUG   = "DEBUG"
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
	FATAL   = "FATAL"
)

const (
	INT_DEBUG   = 1
	INT_INFO    = 2
	INT_WARNING = 3
	INT_ERROR   = 4
	INT_FATAL   = 5
)

var levelMap = map[string]int{
	DEBUG:   INT_DEBUG,
	INFO:    INT_INFO,
	WARNING: INT_WARNING,
	ERROR:   INT_ERROR,
	FATAL:   INT_FATAL,
}

type Logger struct {
	*log.Logger
	level int
}

func NewLogger(out io.Writer, prefix string, flag int, level string) *Logger {
	if _, ok := levelMap[level]; !ok {
		return nil
	}
	return &Logger{log.New(out, prefix, flag), levelMap[level]}
}

func (l *Logger) SetOutput(out io.Writer) {
	if out == nil {
		return
	}
	l.Logger.SetOutput(out)
}

func (l *Logger) SetLevel(level string) {
	if ilvl, ok := levelMap[level]; ok {
		l.level = ilvl
	}
}

func (l *Logger) GetLevel() int {
	return l.level
}

func (l *Logger) GetLevelString() string {
	switch l.level {
	case INT_DEBUG:
		return DEBUG
	case INT_INFO:
		return INFO
	case INT_WARNING:
		return WARNING
	case INT_ERROR:
		return ERROR
	case INT_FATAL:
		return FATAL
	}

	return "WRONG_LEVEL"
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.level > INT_DEBUG {
		return
	}
	if len(v) == 0 {
		l.Println("[DEBUG] " + format)
	} else {
		l.Printf("[DEBUG] "+format, v...)
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level > INT_INFO {
		return
	}
	if len(v) == 0 {
		l.Println("[INFO] " + format)
	} else {
		l.Printf("[INFO] "+format, v...)
	}
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	if l.level > INT_WARNING {
		return
	}
	if len(v) == 0 {
		l.Println("[WARNING] " + format)
	} else {
		l.Printf("[WARNING] "+format, v...)
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.level > INT_ERROR {
		return
	}
	if len(v) == 0 {
		l.Println("[ERROR] " + format)
	} else {
		l.Printf("[ERROR] "+format, v...)
	}
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l.level > INT_FATAL {
		return
	}
	if len(v) == 0 {
		l.Println("[FATAL] " + format)
	} else {
		l.Printf("[FATAL] "+format, v...)
	}
}

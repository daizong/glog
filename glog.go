package glog

import (
	"fmt"
	"io"
	"log"
)

const (
	// DEBUG level
	DEBUG = "DEBUG"
	// INFO level
	INFO = "INFO"
	// WARNING level
	WARNING = "WARNING"
	// ERROR level
	ERROR = "ERROR"
	// FATAL level
	FATAL = "FATAL"
)

const (
	debugFlag   = 1
	infoFlag    = 2
	warningFlag = 3
	errorFlag   = 4
	fatalFlag   = 5
)

var levelMap = map[string]int{
	DEBUG:   debugFlag,
	INFO:    infoFlag,
	WARNING: warningFlag,
	ERROR:   errorFlag,
	FATAL:   fatalFlag,
}

var calldepth = 2

// Logger logger
type Logger struct {
	*log.Logger
	level int
}

// NewLogger return new logger
func NewLogger(out io.Writer, prefix string, flag int, level string) *Logger {
	if _, ok := levelMap[level]; !ok {
		return nil
	}
	return &Logger{log.New(out, prefix, flag), levelMap[level]}
}

// SetOutput set new output
func (l *Logger) SetOutput(out io.Writer) {
	if out == nil {
		return
	}
	l.Logger.SetOutput(out)
}

// SetLevel set new log level
func (l *Logger) SetLevel(level string) {
	if ilvl, ok := levelMap[level]; ok {
		l.level = ilvl
	}
}

// SetCallDepth set new calldepth
func (l *Logger) SetCallDepth(depth int) {
	if depth > 0 {
		calldepth = depth
	}
}

// GetLevel return log level int
func (l *Logger) GetLevel() int {
	return l.level
}

// GetLevelString get log level string
func (l *Logger) GetLevelString() string {
	switch l.level {
	case debugFlag:
		return DEBUG
	case infoFlag:
		return INFO
	case warningFlag:
		return WARNING
	case errorFlag:
		return ERROR
	case fatalFlag:
		return FATAL
	}

	return "WRONG_LEVEL"
}

// Debugf print debug log
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.level > debugFlag {
		return
	}
	if len(v) == 0 {
		l.Output(calldepth, fmt.Sprintln("[DEBUG] "+format))
	} else {
		l.Output(calldepth, fmt.Sprintf("[DEBUG] "+format, v...))
	}
}

// Infof print info log
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level > infoFlag {
		return
	}
	if len(v) == 0 {
		l.Output(calldepth, fmt.Sprintln("[INFO] "+format))
	} else {
		l.Output(calldepth, fmt.Sprintf("[INFO] "+format, v...))
	}
}

// Warningf print warning log
func (l *Logger) Warningf(format string, v ...interface{}) {
	if l.level > warningFlag {
		return
	}
	if len(v) == 0 {
		l.Output(calldepth, fmt.Sprintln("[WARNING] "+format))
	} else {
		l.Output(calldepth, fmt.Sprintf("[WARNING] "+format, v...))
	}
}

// Errorf print error log
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.level > errorFlag {
		return
	}
	if len(v) == 0 {
		l.Output(calldepth, fmt.Sprintln("[ERROR] "+format))
	} else {
		l.Output(calldepth, fmt.Sprintf("[ERROR] "+format, v...))
	}
}

// Fatalf print fatal log
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l.level > fatalFlag {
		return
	}
	if len(v) == 0 {
		l.Output(calldepth, fmt.Sprintln("[FATAL] "+format))
	} else {
		l.Output(calldepth, fmt.Sprintf("[FATAL] "+format, v...))
	}
}

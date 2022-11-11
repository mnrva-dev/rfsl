package rfsl

import (
	"fmt"
	"time"
)

// FileName is the path/name of the file in which logs will be stored (i.e.
// path/to/logs/log.txt).
//
// TimeFormat is the format in which time will be printed in the log. Pass an
// empty string to use the default format, which is 2006-01-02 15:04:05.0000.
//
// MaxSize is the maximum size of each log file before the logger moves on
// to a new file, in bytes. Set to 0 to use only one log file.
// Please note that the log files should not be assumed to be
// at most maxsize bytes large, as they can surpass that size before the logger
// realizes it must move on to a new file.
//
// Level is the logging level. Set it with LOG_LVL_ALL, LOG_LVL_DEBUG,
// LOG_LVL_WARN, etc.
func Config(FileName, TimeFormat string, MaxSize int64, level LogLevel) error {
	err := startrfsl(FileName)
	rfsl.size = MaxSize
	rfsl.level = level
	if TimeFormat == "" {
		rfsl.timeformat = "2006-01-02 15:04:05.0000"
	} else {
		rfsl.timeformat = TimeFormat
	}
	return err
}

// Write a trace to the log file.
//
// Format: '2006-01-02 15:04:05.0000 TRACE	↦ This is a trace log example'
func Trace(v ...any) {
	if rfsl.level <= LOG_LVL_TRACE {
		t := time.Now()
		rfsl.write([]byte(fmt.Sprintf("%s TRACE\t↦ %s\n", t.Format(rfsl.timeformat), fmt.Sprint(v...))))
	}
}

// Write debugging information to the log file.
//
// Format: '2006-01-02 15:04:05.0000 DEBUG	↦ This information is useful for debugging'
func Debug(v ...any) {
	if rfsl.level <= LOG_LVL_DEBUG {
		t := time.Now()
		rfsl.write([]byte(fmt.Sprintf("%s DEBUG\t↦ %s\n", t.Format(rfsl.timeformat), fmt.Sprint(v...))))
	}
}

// Write info to the log file.
//
// Format: '2006-01-02 15:04:05.0000 INFO	↦ This is an info log example'
func Info(v ...any) {
	if rfsl.level <= LOG_LVL_INFO {
		t := time.Now()
		rfsl.write([]byte(fmt.Sprintf("%s INFO\t↦ %s\n", t.Format(rfsl.timeformat), fmt.Sprint(v...))))
	}
}

// Write a warning to the log file.
//
// Format: '2006-01-02 15:04:05.0000 WARN	↦ This information is useful for debugging'
func Warning(v ...any) {
	if rfsl.level <= LOG_LVL_WARN {
		t := time.Now()
		rfsl.write([]byte(fmt.Sprintf("%s WARN\t↦ %s\n", t.Format(rfsl.timeformat), fmt.Sprint(v...))))
	}
}

// Write an error to the log file.
//
// Format: '2006-01-02 15:04:05.0000 ERROR	↦ There is an error here'
func Error(v ...any) {
	if rfsl.level <= LOG_LVL_ERROR {
		t := time.Now()
		rfsl.write([]byte(fmt.Sprintf("%s ERROR\t↦ %s\n", t.Format(rfsl.timeformat), fmt.Sprint(v...))))
	}
}

// Write a panic to the log file.
//
// Format: '2006-01-02 15:04:05.0000 PANIC	↦ This is a panic!'
func Panic(v ...any) {
	if rfsl.level <= LOG_LVL_PANIC {
		t := time.Now()
		rfsl.write([]byte(fmt.Sprintf("%s PANIC\t↦ %s\n", t.Format(rfsl.timeformat), fmt.Sprint(v...))))
	}
}

// Write a fatal error to the log file.
//
// Format: '2006-01-02 15:04:05.0000 FATAL	↦ This is a fatal error!'
func Fatal(v ...any) {
	t := time.Now()
	rfsl.write([]byte(fmt.Sprintf("%s FATAL\t↦ %s\n", t.Format(rfsl.timeformat), fmt.Sprint(v...))))
}

// Trace, with printf style syntax.
func Tracef(format string, v ...any) {
	if rfsl.level <= LOG_LVL_TRACE {
		t := time.Now()
		rfsl.write([]byte(fmt.Sprintf("%s TRACE\t↦ %s\n", t.Format(rfsl.timeformat), fmt.Sprintf(format, v...))))
	}
}

// Debug, with printf style syntax.
func Debugf(format string, v ...any) {
	if rfsl.level <= LOG_LVL_DEBUG {
		t := time.Now()
		rfsl.write([]byte(fmt.Sprintf("%s DEBUG\t↦ %s\n", t.Format(rfsl.timeformat), fmt.Sprintf(format, v...))))
	}
}

// Info, with printf style syntax.
func Infof(format string, v ...any) {
	if rfsl.level <= LOG_LVL_INFO {
		t := time.Now()
		rfsl.write([]byte(fmt.Sprintf("%s INFO\t↦ %s\n", t.Format(rfsl.timeformat), fmt.Sprintf(format, v...))))
	}
}

// Warning, with printf style syntax.
func Warningf(format string, v ...any) {
	if rfsl.level <= LOG_LVL_WARN {
		t := time.Now()
		rfsl.write([]byte(fmt.Sprintf("%s WARN\t↦ %s\n", t.Format(rfsl.timeformat), fmt.Sprintf(format, v...))))
	}
}

// Error, with printf style syntax.
func Errorf(format string, v ...any) {
	if rfsl.level <= LOG_LVL_ERROR {
		t := time.Now()
		rfsl.write([]byte(fmt.Sprintf("%s ERROR\t↦ %s\n", t.Format(rfsl.timeformat), fmt.Sprintf(format, v...))))
	}
}

// Panic, with printf style syntax.
func Panicf(format string, v ...any) {
	if rfsl.level <= LOG_LVL_PANIC {
		t := time.Now()
		rfsl.write([]byte(fmt.Sprintf("%s PANIC\t↦ %s\n", t.Format(rfsl.timeformat), fmt.Sprintf(format, v...))))
	}
}

// Fatal, with printf style syntax.
func Fatalf(format string, v ...any) {
	t := time.Now()
	rfsl.write([]byte(fmt.Sprintf("%s FATAL\t↦ %s\n", t.Format(rfsl.timeformat), fmt.Sprintf(format, v...))))
}

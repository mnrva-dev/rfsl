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
func Config(FileName, TimeFormat string, MaxSize int64) error {
	err := startRful(FileName)
	rful.size = MaxSize
	if TimeFormat == "" {
		rful.timeformat = "2006-01-02 15:04:05.0000"
	} else {
		rful.timeformat = TimeFormat
	}
	return err
}

// Write a trace to the log file.
//
// Format: '2006-01-02 15:04:05.0000	This is a trace log example'
func Trace(v ...any) {
	t := time.Now()
	rful.write([]byte(fmt.Sprintf("%s\t%s\n", t.Format(rful.timeformat), fmt.Sprint(v...))))
}

// Write info to the log file.
//
// Format: '2006-01-02 15:04:05.0000 INFO:	This is an info log example'
func Info(v ...any) {
	t := time.Now()
	rful.write([]byte(fmt.Sprintf("%s INFO:\t%s\n", t.Format(rful.timeformat), fmt.Sprint(v...))))
}

// Write debugging information to the log file.
//
// Format: '2006-01-02 15:04:05.0000 DEBUG:	This information is useful for debugging'
func Debug(v ...any) {
	t := time.Now()
	rful.write([]byte(fmt.Sprintf("%s DEBUG:\t%s\n", t.Format(rful.timeformat), fmt.Sprint(v...))))
}

// Write a warning to the log file.
//
// Format: '2006-01-02 15:04:05.0000 WARNING:	This information is useful for debugging'
func Warning(v ...any) {
	t := time.Now()
	rful.write([]byte(fmt.Sprintf("%s WARNING:\t%s\n", t.Format(rful.timeformat), fmt.Sprint(v...))))
}

// Write an error to the log file.
//
// Format: '2006-01-02 15:04:05.0000 *ERROR:	There is an error here'
func Error(v ...any) {
	t := time.Now()
	rful.write([]byte(fmt.Sprintf("%s *ERROR:\t%s\n", t.Format(rful.timeformat), fmt.Sprint(v...))))
}

// Write a panic to the log file.
//
// Format: '2006-01-02 15:04:05.0000 **PANIC**	This is a panic!'
func Panic(v ...any) {
	t := time.Now()
	rful.write([]byte(fmt.Sprintf("%s **PANIC**\t%s\n", t.Format(rful.timeformat), fmt.Sprint(v...))))
}

// Write a trace to the log file using printf-style formatting.
//
// Format: '2006-01-02 15:04:05.0000	This is a trace log example'
func Tracef(format string, v ...any) {
	t := time.Now()
	rful.write([]byte(fmt.Sprintf("%s\t%s\n", t.Format(rful.timeformat), fmt.Sprintf(format, v...))))
}

// Write info to the log file using printf-style formatting.
//
// Format: '2006-01-02 15:04:05.0000 INFO:	This is an info log example'
func Infof(format string, v ...any) {
	t := time.Now()
	rful.write([]byte(fmt.Sprintf("%s INFO:\t%s\n", t.Format(rful.timeformat), fmt.Sprintf(format, v...))))
}

// Write debugging information to the log file using printf-style formatting.
//
// Format: '2006-01-02 15:04:05.0000 DEBUG:	This information is useful for debugging'
func Debugf(format string, v ...any) {
	t := time.Now()
	rful.write([]byte(fmt.Sprintf("%s DEBUG:\t%s\n", t.Format(rful.timeformat), fmt.Sprintf(format, v...))))
}

// Write a warning to the log file using printf-style formatting.
//
// Format: '2006-01-02 15:04:05.0000 WARNING:	This information is useful for debugging'
func Warningf(format string, v ...any) {
	t := time.Now()
	rful.write([]byte(fmt.Sprintf("%s WARNING:\t%s\n", t.Format(rful.timeformat), fmt.Sprintf(format, v...))))
}

// Write an error to the log file using printf-style formatting.
//
// Format: '2006-01-02 15:04:05.0000 *ERROR:	There is an error here'
func Errorf(format string, v ...any) {
	t := time.Now()
	rful.write([]byte(fmt.Sprintf("%s *ERROR:\t%s\n", t.Format(rful.timeformat), fmt.Sprintf(format, v...))))
}

// Write a panic to the log file using printf-style formatting.
//
// Format: '2006-01-02 15:04:05.0000 **PANIC**	This is a panic!'
func Panicf(format string, v ...any) {
	t := time.Now()
	rful.write([]byte(fmt.Sprintf("%s **PANIC**\t%s\n", t.Format(rful.timeformat), fmt.Sprintf(format, v...))))
}

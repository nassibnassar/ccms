package log

import (
	"fmt"
	"io"
	glog "log"
	"os"
	"sync"
	"time"

	"github.com/indexdata/ccms/cmd/ccms/osutil"
)

type Log struct {
	log      *glog.Logger
	logDebug bool
	logTrace bool
}

var std Log
var once sync.Once

func Init(out io.Writer, logDebug bool, logTrace bool) {
	if out != nil {
		once.Do(func() {
			std = Log{
				log:      glog.New(out, "", 0),
				logDebug: logDebug,
				logTrace: logTrace,
			}
		})
	}
}

func Fatal(format string, args ...interface{}) {
	printf("FATAL", format, args...)
}

func Error(format string, args ...interface{}) {
	printf("ERROR", format, args...)
}

func Warning(format string, args ...interface{}) {
	printf("WARNING", format, args...)
}

func Info(format string, args ...interface{}) {
	printf("INFO", format, args...)
}

func Debug(format string, args ...interface{}) {
	if !std.logDebug && !std.logTrace {
		return
	}
	printf("DEBUG", format, args...)
}

func Trace(format string, args ...interface{}) {
	if !std.logTrace {
		return
	}
	printf("TRACE", format, args...)
}

func Detail(format string, args ...interface{}) {
	printf("DETAIL", format, args...)
}

func IsLevelTrace() bool {
	return std.logTrace
}

func printf(level string, format string, args ...interface{}) {
	var msg = fmt.Sprintf(format, args...)
	var n = time.Now().UTC()
	var now = n.Format("2006-01-02 15:04:05 MST")
	std.log.Printf("%s  %s  %s", now, level+":", msg)
}

func OpenLogFile(logfile string) (*os.File, error) {
	var f *os.File
	var err error
	if f, err = os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, osutil.ModePermRW); err != nil {
		return nil, err
	}
	return f, nil
}

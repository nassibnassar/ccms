package eout

import (
	"fmt"
	"os"
)

var EnableVerbose bool
var EnableTrace bool

func Init(program string) {
	prog = program
}

func Error(format string, v ...interface{}) {
	message(format, v...)
}

func Warning(format string, v ...interface{}) {
	message(format, v...)
}

func Info(format string, v ...interface{}) {
	message(format, v...)
}

func Verbose(format string, v ...interface{}) {
	if !EnableVerbose && !EnableTrace {
		return
	}
	message(format, v...)
}

func Trace(format string, v ...interface{}) {
	if !EnableTrace {
		return
	}
	message(format, v...)
}

func message(format string, v ...interface{}) {
	_, _ = fmt.Fprintf(std, format+"\n", v...)
}

var std *os.File = os.Stderr
var prog string

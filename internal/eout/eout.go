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
	locus()
	//_, _ = ErrorColor.Fprint(std, "error: ")
	message(format, v...)
}

func Warning(format string, v ...interface{}) {
	locus()
	//_, _ = WarningColor.Fprint(std, "warning: ")
	message(format, v...)
}

func Info(format string, v ...interface{}) {
	locus()
	message(format, v...)
}

func Verbose(format string, v ...interface{}) {
	if !EnableVerbose && !EnableTrace {
		return
	}
	locus()
	message(format, v...)
}

func Trace(format string, v ...interface{}) {
	if !EnableTrace {
		return
	}
	message(format, v...)
}

func locus() {
	if !interactive {
		_, _ = LocusColor.Fprint(std, fmt.Sprintf("%s: ", prog))
	}
}

func message(format string, v ...interface{}) {
	_, _ = fmt.Fprintf(std, format+"\n", v...)
}

func Interactive() {
	interactive = true
}

var std *os.File = os.Stderr
var prog string
var interactive bool

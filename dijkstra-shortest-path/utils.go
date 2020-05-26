// various utils
package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

var traceOn bool = false
var debugOn bool = false

func EnableTrace() {
	traceOn = true
}

func EnableDebug() {
	debugOn = true
}

func Trace() {
	if traceOn {
		_, file, line, ok := runtime.Caller(1 /* the '1' is to skip trace() function itself */)
		if ok {
			idx := strings.LastIndex(file, "/")
			fmt.Printf("TRACE: %s %d\n", file[idx+1:], line)
		}
	}
}

func PanicIfError(err error) {
	if err != nil {
		Trace()
		panic(err)
	}
}

func PanicIfFalse(res bool, msg string) {
	if !res {
		Trace()
		panic(msg)
	}
}

func PrintDebug(format string, args ...interface{}) {
	if debugOn {
		fmt.Printf("DEBUG: "+format+"\n", args...)
	}
}

func PrintExecTime(start time.Time, name string) {
	if debugOn {
		PrintDebug("%s took %s", name, time.Since(start))
	}
}

// end of file

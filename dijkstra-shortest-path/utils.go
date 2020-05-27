// various utils
package main

import (
	"bufio"
	"fmt"
	"os"
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

func PanicIf(cond bool, format string, args ...interface{}) {
	if cond {
		Trace()
		panic(fmt.Sprintf(format, args...))
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

// Read 'fname' file line by line.
// Process incoming data with 'handler' function.
func readFilePerLine(fname string, handler func(line string) error) (e error) {

	Trace()

	file, err := os.Open(fname)
	defer func() {
		file.Close()
		r := recover()
		err, ok := r.(error) // typecast 'r' to 'error'
		if ok {
			e = err // assign 'e' (named function result) to captured error
		}
	}()
	PanicIfError(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		PanicIfError(handler(scanner.Text()))
	}

	PanicIfError(scanner.Err())

	Trace()
	return nil
}

// end of file

// Dijkstra shortest path

package main

import (
	"bufio"
	//"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

var traceOn bool = true
var debugOn bool = true

func trace() {
	if traceOn {
		_, file, line, ok := runtime.Caller(1 /* the '1' is to skip trace() function itself */)
		if ok {
			idx := strings.LastIndex(file, "/")
			fmt.Printf("TRACE: %s %d\n", file[idx+1:], line)
		}
	}
}

func panicIfError(err error) {
	if err != nil {
		trace()
		panic(err)
	}
}

func panicIfFalse(res bool, msg string) {
	if !res {
		trace()
		panic(msg)
	}
}

func printDebug(format string, args ...interface{}) {
	if debugOn {
		fmt.Printf("DEBUG: "+format+"\n", args...)
	}
}

func printExecTime(start time.Time, name string) {
	if debugOn {
		printDebug("%s took %s", name, time.Since(start))
	}
}

type edge struct {
	from int
	to   int
	cost int
}

// Read 'fname' file line by line.
// Process incoming data with 'handler' function.
func readFilePerLine(fname string, handler func(line string) error) (e error) {

	trace()

	file, err := os.Open(fname)
	defer func() {
		file.Close()
		r := recover()
		err, ok := r.(error) // typecast 'r' to 'error'
		if ok {
			e = err // assign 'e' (named function result) to captured error
		}
	}()
	panicIfError(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		panicIfError(handler(scanner.Text()))
	}

	panicIfError(scanner.Err())

	trace()
	return nil
}

// Read graph edges from text file.
// Return number of nodes, array of edges and error.
// First line of the file must contain 'numberOfNodes numberOfEdges'.
// All the rest lines must be edges of 'from to cost' format.
// Format defined by Algo2 Coursera course in Dijkstra shortest path task.
func readGraph(fname string) (nodeCount int, edges []edge, err error) {

	start := time.Now()

	trace()
	lineNo, edgeCount := 0, 0

	nodeCount = 0

	err = readFilePerLine(fname,
		func(line string) error {
			lineNo++
			if lineNo == 1 {
				// parse 'nodeCount edgeCount'
				n, e := fmt.Sscanf(line, "%d %d", &nodeCount, &edgeCount)
				panicIfFalse(n == 2 && e == nil,
					fmt.Sprintf("failed to read %s, line %d: bad format", fname, lineNo))
			} else {
				// parse 'from to cost'
				var from, to, cost int
				n, e := fmt.Sscanf(line, "%d %d %d", &from, &to, &cost)
				panicIfFalse(n == 3 && e == nil,
					fmt.Sprintf("failed to read %s, line %d: bad format", fname, lineNo))
				c1 := cap(edges)
				edges = append(edges, edge{from, to, cost})
				c2 := cap(edges)
				if c2 != c1 {
					printDebug("readGraph() resize: %d->%d (+%d)", c1, c2, c2-c1)
				}
			}
			return nil
		})
	if err != nil {
		return 0, nil, err
	}

	trace()

	printExecTime(start, "readGraph()")

	return nodeCount, edges, nil
}

func main() {
	traceOn = false
	debugOn = false

	for n, arg := range os.Args {
		if n > 0 {
			switch arg {
			case "-d":
				debugOn = true
			case "-t":
				traceOn = true
			}
		}
	}

	trace()

	nodeCount, edges, err := readGraph("large.txt")
	panicIfError(err)

	fmt.Printf("nodes %d, edges %d, %v...\n", nodeCount, len(edges), edges[:3])

	trace()
}

// end of file

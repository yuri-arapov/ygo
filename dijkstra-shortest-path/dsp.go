// Dijkstra shortest path

package main

import (
	"bufio"
	//"errors"
	"fmt"
	"os"
	"time"
)

type edge struct {
	from int
	to   int
	cost int
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

// Read graph edges from text file.
// Return number of nodes, array of edges and error.
// First line of the file must contain 'numberOfNodes numberOfEdges'.
// All the rest lines must be edges of 'from to cost' format.
// Format defined by Algo2 Coursera course in Dijkstra shortest path task.
func readGraph(fname string) (nodeCount int, edges []edge, err error) {

	start := time.Now()

	Trace()
	lineNo, edgeCount := 0, 0

	nodeCount = 0

	err = readFilePerLine(fname,
		func(line string) error {
			lineNo++
			if lineNo == 1 {
				// parse 'nodeCount edgeCount'
				n, e := fmt.Sscanf(line, "%d %d", &nodeCount, &edgeCount)
				PanicIfFalse(n == 2 && e == nil,
					fmt.Sprintf("failed to read %s, line %d: bad format", fname, lineNo))
			} else {
				// parse 'from to cost'
				var from, to, cost int
				n, e := fmt.Sscanf(line, "%d %d %d", &from, &to, &cost)
				PanicIfFalse(n == 3 && e == nil,
					fmt.Sprintf("failed to read %s, line %d: bad format", fname, lineNo))
				c1 := cap(edges)
				edges = append(edges, edge{from, to, cost})
				c2 := cap(edges)
				if c2 != c1 {
					PrintDebug("readGraph() resize: %d->%d (+%d)", c1, c2, c2-c1)
				}
			}
			return nil
		})
	if err != nil {
		return 0, nil, err
	}

	Trace()

	PrintExecTime(start, "readGraph()")

	return nodeCount, edges, nil
}

func main() {
	for n, arg := range os.Args {
		if n > 0 {
			switch arg {
			case "-d":
				EnableDebug()
			case "-t":
				EnableTrace()
			}
		}
	}

	Trace()

	const fname string = "large.txt"
	nodeCount, edges, err := readGraph(fname)
	PanicIfError(err)

	fmt.Printf("%s: nodes %d, edges %d, %v...\n", fname, nodeCount, len(edges), edges[:3])

	Trace()
}

// end of file

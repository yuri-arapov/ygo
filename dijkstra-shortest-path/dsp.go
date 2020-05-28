// Dijkstra shortest path

package main

import (
	"fmt"
	"os"
	"time"
	"yheap"
)

// edge
type edge struct {
	from int
	to   int
	cost int
}

// rib
type rib struct {
	to   int
	cost int
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
				PanicIf(n != 2 || e != nil, "failed to read %s, line %d: bad format", fname, lineNo)
			} else {
				// parse 'from to cost'
				var from, to, cost int
				n, e := fmt.Sscanf(line, "%d %d %d", &from, &to, &cost)
				PanicIf(n != 3 || e != nil, "failed to read %s, line %d: bad format", fname, lineNo)
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

// return minimum throughout arguments
func min(x int, rest ...int) int {
	for _, r := range rest {
		if r < x {
			x = r
		}
	}
	return x
}

// return maximum throughout arguments
func max(x int, rest ...int) int {
	for _, r := range rest {
		if r > x {
			x = r
		}
	}
	return x
}

// return largest node in the graph
func maxGraphNode(g []edge) int {
	n := 0
	for _, e := range g {
		n = max(n, e.from, e.to)
	}
	return n
}

// turn list of edges into adjacency list
// i-th element of adjacency list is an array of ribs the i-th node connected to
func edgesToAdjList(g []edge) [][]rib {
	n := maxGraphNode(g) // number of nodes
	res := make([][]rib, n+1)
	for _, e := range g {
		res[e.from] = append(res[e.from], rib{e.to, e.cost})
	}
	return res
}

// Dijkstra shortest path algorithm
func dijrstraShortestPath(g []edge, s int) []int {
	// number of nodes
	n := maxGraphNode(g)

	// number of edges
	m := len(g)

	// less function
	less := func(x, y int) bool { return x < y }

	// heap to speed up algorithm
	h := yheap.MakeKHeap(m, less)

	// adjacensy list. the ag[i] is a list of ribs the i-th node connected to
	ag := edgesToAdjList(g)

	// tracking of minimum path
	a := makeArray(n+1, -1)

	explored := func(node int) bool { return a[node] != -1 }

	// counter of unresolved nodes
	unresolvedCount := n

	// resolve node: put it into solution set and update outgoing edges
	resolveNode := func(node int, distance int) {
		unresolvedCount--

		// save distance from starting node 's' to given 'node'
		a[node] = distance

		// for each outgoing edge of given 'node'
		for _, r := range ag[node] {
			head := r.to
			if !explored(head) {
				if h.Contains(head) {
					h.Update(head, min(h.GetKey(head), distance+r.cost))
				} else {
					h.Push(head, distance+r.cost)
				}
			}
		}
	}

	// put source node into solution
	resolveNode(s, 0)

	// resolve all the rest nodes taking the nearest one from the heap
	for unresolvedCount > 0 {
		node, cost := h.Pop()
		resolveNode(node, cost)
	}

	return a
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

	start := time.Now()
	res := dijrstraShortestPath(edges, 1)
	PrintExecTime(start, "Dijkstra shortest path algorithm")
	fmt.Printf("%v...\n", res[:10])

	Trace()
}

// end of file

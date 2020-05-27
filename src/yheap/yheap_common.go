package yheap

import (
	"fmt"
	"math"
)

type Less func(a, b int) bool

func log2(x int) float64 {
	return math.Log(float64(x)) / math.Log(2.0)
}

func parent(pos int) int { return (pos - 1) / 2 }
func left(pos int) int   { return pos*2 + 1 }
func right(pos int) int  { return pos*2 + 2 }

func bTreeHeight(size int) int {
	return int(math.Round(math.Ceil(log2(size + 1))))
}

func panicIf(cond bool, format string, args ...interface{}) {
	if cond {
		panicPrintf(format, args...)
	}
}

func panicPrintf(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
}

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
	"yheap"
)

func shuffle(data *[]int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	size := len(*data)
	for i := 0; i < size; i++ {
		x := r.Intn(size)
		y := r.Intn(size)
		(*data)[x], (*data)[y] = (*data)[y], (*data)[x]
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	const size = 1000000
	const show = 10

	fmt.Println("Heap test program, size", size)

	config := []struct {
		what string
		less yheap.Less
	}{
		{"min heap", func(a, b int) bool { return a < b }},
		{"max heap", func(a, b int) bool { return a > b }}}

	for _, cfg := range config {
		h := yheap.MakeKHeap(cfg.less, size)
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}

		shuffle(&data)
		fmt.Println("randomized", data[:show], "...")

		sorted := make([]int, size)
		copy(sorted, data)
		sort.Slice(sorted,
			func(i, j int) bool { return cfg.less(sorted[i], sorted[j]) })
		fmt.Println("sorted    ", sorted[:show], "...")

		for _, x := range data {
			h.Push(x, x)
		}
		fmt.Println("heap height", h.Height())

		sortedByHeap := make([]int, size)
		for i, _ := range data {
			sortedByHeap[i], _ = h.Pop()
		}
		fmt.Println("from heap ", sortedByHeap[:show], "...")

		///foo := []int{0, 0, 0}
		if equal(sorted, sortedByHeap) {
			fmt.Println(cfg.what, "OK")
		} else {
			fmt.Println(cfg.what, "sort failed")
		}
	}
}

// end of file

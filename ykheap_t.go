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
	const size = 100
	const show = 10

	fmt.Println("Heap test program, size", size)

	lessThan := func(a, b int) bool { return a < b }
	greaterThan := func(a, b int) bool { return a > b }

	config := []struct {
		what string
		less yheap.Less
	}{{"min heap", lessThan}, {"max heap", greaterThan}}

	for _, cfg := range config {
		h := yheap.MakeKHeap(size, cfg.less)
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
		fmt.Printf("**********************************************\n")
	}

	fmt.Printf("**********************************************\n")
	const s = 14
	h := yheap.MakeKHeap(s, lessThan)
	data := make([]int, s)
	for i := range data {
		data[i] = i
	}
	shuffle(&data)
	for i, n := range data {
		h.Push(i, n)
	}
	node := s / 2
	fmt.Printf("updating node %d\n", node)
	PrintHeap(h)
	for x := 0; x < s; x++ {
		key := h.GetKey(node)
		newKey := key - 1
		fmt.Printf("node %d: %d->%d\n", node, key, newKey)
		h.Update(node, newKey)
		PrintHeap(h)
		fmt.Printf("-----\n")
	}

	/*
		fmt.Printf("**********************************************\n")
		h2 := yheap.MakeKHeap(s, lessThan)
		h3 := yheap.MakeKHeap(s, lessThan)
		for i, n := range data {
			h2.Push(i, n)
			h3.Push(i, n)
		}
		for x := 0; x < s; x++ {
			key := h2.GetKey(node)
			if key != h3.GetKey(node) {
				fmt.Printf("!!! keys mismatch node=%d!!!\n", node)
			}
			h2.Update(node, key-1)
			h3.UpdateInplace(node, key-1)
			p2, p3 := h2.GetPos(node), h3.GetPos(node)
			if p2 != p3 {
				fmt.Printf("pos mismatch: node=%d pos2=%d pos3=%d\n", node, p2, p3)
			} else {
				fmt.Printf("ok\n")
			}
		}
	*/
}

// end of file

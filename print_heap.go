package main

import (
	"fmt"
	"yheap"
)

func pwr(n, p int) int {
	if p == 0 {
		return 1
	} else {
		return n * pwr(n, p-1)
	}
}

func max(n int, rest ...int) int {
	for _, r := range rest {
		if r > n {
			n = r
		}
	}
	return n
}

func buff(s string, width int) string {
	for left := true; len(s) < width; left = !left {
		if left {
			s = " " + s
		} else {
			s = s + " "
		}
	}
	return s
}

// Print heap as binary tree.
func PrintHeap(h yheap.KHeap) {
	if h.Count() == 0 {
		return
	}

	const gap string = "   "

	elementWidth := 0

	// pass 1: stringify heap elements
	elements := make([]string, 0)
	for i := 0; i < h.Count(); i++ {
		node, key := h.At(i)
		s := fmt.Sprintf("%d,%d", node, key)
		elementWidth = max(elementWidth, len(s))
		elements = append(elements, s)
	}

	// number of elements of the bottom line of the tree
	lastLineCount := pwr(2, h.Height()-1)

	// width of the last line (elements and gap between them)
	lastLineWidth := lastLineCount*elementWidth + (lastLineCount-1)*len(gap)

	// for each line of the tree
	for line := 0; line < h.Height(); line++ {
		// number of elements per line
		count := pwr(2, line)

		// width of the whitespace of the line
		whiteSpace := lastLineWidth - count*elementWidth

		// width of separator between elements
		separator := whiteSpace / count

		// width of the first element offset
		offset := separator / 2

		// for each element of the line
		for pos := count - 1; count > 0; count-- {
			if pos < len(elements) {
				if pos == count-1 {
					// first element in the line
					fmt.Printf("%s", buff("", offset))
				} else {
					fmt.Printf("%s", buff("", separator))
				}
				fmt.Printf("%s", buff(elements[pos], elementWidth))
			}
			pos++
		}
		fmt.Printf("\n")
	}
}

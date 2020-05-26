// KeyHeap
//
// Keeps track of min/max key-node pairs.

package ykheap

import "math"
import "fmt"

type Less func(a, b int) bool

type KHeap struct {
	less    Less  // '<' operator
	size    int   // size of the heap
	count   int   // number of elements in the heap
	nodeKey []int // nodeKey[n] is a key of node n
	nodePos []int // nodePos[n] is a pos of node n
	posNode []int // posNode[p] is a node value of pos p
}

const badPos int = -1

func MakeHeap(less Less, size int) KHeap {
	nodeKey := make([]int, size)
	nodePos := make([]int, size)
	for i, _ := range nodePos {
		nodePos[i] = badPos
	}
	posNode := make([]int, size)
	return KHeap{less, size, 0, nodeKey, nodePos, posNode}
}

func (h *KHeap) Size() int   { return h.size }
func (h *KHeap) Count() int  { return h.count }
func (h *KHeap) Height() int { return int(math.Round(math.Ceil(log2(h.count + 1)))) }

func (h *KHeap) Contains(node int) bool {
	return node >= 0 && node < h.count && h.nodePos[node] != badPos
}

func (h *KHeap) Push(node, key int) {
	if h.count == h.size {
		panic("Push: KHeap is full")
	}
	if h.Contains(node) {
		panic(fmt.Sprintf("Push: node already in the heap (%d)", node))
	}
	pos := h.count
	h.nodeKey[node] = key
	h.nodePos[node] = pos
	h.posNode[pos] = node
	h.count++
	h.heapifyUp(pos)
}

func (h *KHeap) Pop() (node, key int) {
	if h.count == 0 {
		panic("Pop: KHeap is empty")
	}
	node = h.posNode[0]
	key = h.nodeKey[node]
	h.nodeKey[node] = -1
	h.swap(0, h.count-1)
	h.count--
	h.heapifyDown(0)
	return
}

func (h *KHeap) Top() (node, key int) {
	if h.count == 0 {
		panic("Top: KHeap is empty")
	}
	return h.posNode[0], h.nodeKey[h.posNode[0]]
}

func (h *KHeap) parent(pos int) int { return (pos - 1) / 2 }
func (h *KHeap) left(pos int) int   { return pos*2 + 1 }
func (h *KHeap) right(pos int) int  { return pos*2 + 2 }

func (h *KHeap) swap(p1, p2 int) {
	n1, n2 := h.posNode[p1], h.posNode[p2]
	h.nodePos[n1] = p2
	h.posNode[p2] = n1
	h.nodePos[n2] = p1
	h.nodePos[p1] = n2
}

func (h *KHeap) posKey(pos int) int {
	return h.nodeKey[h.posNode[pos]]
}

func (h *KHeap) lessByPos(p1, p2 int) bool {
	return h.less(h.posKey(p1), h.posKey(p2))
}

func (h *KHeap) heapifyUp(pos int) {
	if pos > 0 {
		parent := h.parent(pos)
		if h.lessByPos(pos, parent) {
			h.swap(pos, parent)
			h.heapifyUp(parent)
		}
	}
}

func (h *KHeap) heapifyDown(pos int) {
	smaller := pos
	for _, p := range []int{h.left(pos), h.right(pos)} {
		if p < h.count && h.lessByPos(p, smaller) {
			smaller = p
		}
	}
	if smaller != pos {
		h.swap(smaller, pos)
		h.heapifyDown(smaller)
	}
}

func log2(x int) float64 {
	return math.Log(float64(x)) / math.Log(2.0)
}

func init() {
	// do some initialization here
}

// end of file

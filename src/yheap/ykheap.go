// KeyHeap
//
// Keeps track of min/max key-node pairs.
//
// operations:
//   MakeHeap
//   Size
//   Count
//   Height
//   Contains
//   Push
//   Pop
//   Top
//   GetKey
//   Delete
//   Update

package yheap

type KHeap struct {
	less    Less  // '<' operator
	size    int   // size of the heap
	count   int   // number of elements in the heap
	nodeKey []int // nodeKey[n] is a key of node n
	nodePos []int // nodePos[n] is a pos of node n
	posNode []int // posNode[p] is a node value of pos p
}

const (
	badPos  int = -1
	badNode int = -1
)

func MakeKHeap(less Less, size int) KHeap {
	nodeKey := make([]int, size)
	nodePos := make([]int, size)
	posNode := make([]int, size)
	for i, _ := range nodePos {
		nodePos[i] = badPos
		posNode[i] = badNode
	}
	return KHeap{less, size, 0, nodeKey, nodePos, posNode}
}

func (h *KHeap) Size() int   { return h.size }
func (h *KHeap) Count() int  { return h.count }
func (h *KHeap) Height() int { return bTreeHeight(h.count + 1) }

func (h *KHeap) Contains(node int) bool {
	return h.isNodeInRange(node) && h.nodePos[node] != badPos
}

func (h *KHeap) Push(node, key int) {
	panicIf(h.Contains(node), "Push: already in heap: %d", node)
	panicIf(!h.isNodeInRange(node), "Push: out of range: %d", node)
	pos := h.count
	h.nodeKey[node] = key
	h.nodePos[node] = pos
	h.posNode[pos] = node
	h.count++
	h.heapifyUp(pos)
}

func (h *KHeap) Pop() (node, key int) {
	panicIf(h.count == 0, "Top: heap empty")
	node = h.posNode[0]
	key = h.nodeKey[node]
	h.nodePos[node] = badPos
	h.swap(0, h.count-1)
	h.count--
	h.heapifyDown(0)
	return
}

func (h *KHeap) Top() (node, key int) {
	panicIf(h.count == 0, "Top: heap empty")
	return h.posNode[0], h.nodeKey[h.posNode[0]]
}

func (h *KHeap) GetKey(node int) int {
	panicIf(!h.Contains(node), "GetKey: not in heap: %d", node)
	return h.nodeKey[node]
}

func (h *KHeap) Delete(node int) {
	panicIf(!h.Contains(node), "Delete: not in heap: %d", node)
	pos := h.nodePos[node]
	last := h.count - 1
	if pos == last {
		h.count--
		h.nodePos[node] = badPos
	} else {
		h.swap(pos, last)
		h.count--
		h.nodePos[node] = badPos
		h.heapify(pos)
	}
}

func (h *KHeap) Update(node int, newKey int) {
	panicIf(!h.Contains(node), "Update: not in heap: %d", node)
	h.Delete(node)
	h.Push(node, newKey)
}

func (h *KHeap) hasParent(pos int) bool { return pos > 0 }

func (h *KHeap) swap(p1, p2 int) {
	n1, n2 := h.posNode[p1], h.posNode[p2]
	h.nodePos[n1] = p2
	h.posNode[p2] = n1
	h.nodePos[n2] = p1
	h.posNode[p1] = n2
}

func (h *KHeap) posKey(pos int) int        { return h.nodeKey[h.posNode[pos]] }
func (h *KHeap) lessByPos(p1, p2 int) bool { return h.less(h.posKey(p1), h.posKey(p2)) }
func (h *KHeap) isGoodPos(pos int) bool    { return 0 <= pos && pos < h.count }

func (h *KHeap) minOf(p1, p2 int) int {
	switch {
	case !h.isGoodPos(p1):
		return p2
	case !h.isGoodPos(p2):
		return p1
	case h.lessByPos(p1, p2):
		return p1
	default:
		return p2
	}
}

func (h *KHeap) heapifyUp(pos int) {
	if pos > 0 {
		parent := parent(pos)
		if h.lessByPos(pos, parent) {
			h.swap(pos, parent)
			h.heapifyUp(parent)
		}
	}
}

func (h *KHeap) heapifyDown(pos int) {
	smaller := pos
	for _, p := range []int{left(pos), right(pos)} {
		if p < h.count && h.lessByPos(p, smaller) {
			smaller = p
		}
	}
	if smaller != pos {
		h.swap(smaller, pos)
		h.heapifyDown(smaller)
	}
}

func (h *KHeap) heapify(pos int) {
	parent := parent(pos)
	left := left(pos)
	right := right(pos)
	switch {
	case h.hasParent(pos) && h.lessByPos(pos, parent):
		h.swap(pos, parent)
		h.heapify(parent)
	case h.isGoodPos(left) && h.lessByPos(left, pos):
		h.swap(pos, left)
		h.heapify(left)
	case h.isGoodPos(right) && h.lessByPos(right, pos):
		h.swap(pos, right)
		h.heapify(right)
	default:
		// heap property resotred
	}
}

func (h *KHeap) isNodeInRange(node int) bool {
	return 0 <= node && node < h.size
}

func init() {
	// do some initialization here
}

// end of file

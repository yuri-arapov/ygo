// Heap

package yheap

type Less func(a, b int) bool

type Heap struct {
	less  Less  // '<' operator
	size  int   // size of the heap
	count int   // number of elements in the heap
	data  []int // data
}

func MakeHeap(less Less, size int) Heap {
	return Heap{less, size, 0, make([]int, size)}
}

func (h *Heap) Size() int   { return h.size }
func (h *Heap) Count() int  { return h.count }
func (h *Heap) Data() []int { return h.data }

func (h *Heap) Push(i int) {
	if h.count == h.size {
		panic("Push: Heap is full")
	}
	h.data[h.count] = i
	h.count++
	h.heapifyUp(h.count - 1)
}

func (h *Heap) Pop() int {
	if h.count == 0 {
		panic("Pop: Heap is empty")
	}
	res := h.data[0]
	h.swap(0, h.count-1)
	h.count--
	h.heapifyDown(0)
	return res
}

func (h *Heap) Top() int {
	if h.count == 0 {
		panic("Top: Heap is empty")
	}
	return h.data[0]
}

func (h *Heap) parent(pos int) int { return (pos - 1) / 2 }
func (h *Heap) left(pos int) int   { return pos*2 + 1 }
func (h *Heap) right(pos int) int  { return pos*2 + 2 }

func (h *Heap) swap(p1, p2 int) {
	h.data[p1], h.data[p2] = h.data[p2], h.data[p1]
}

func (h *Heap) lessByPos(p1, p2 int) bool {
	return h.less(h.data[p1], h.data[p2])
}

func (h *Heap) heapifyUp(pos int) {
	if pos > 0 {
		parent := h.parent(pos)
		if h.lessByPos(pos, parent) {
			h.swap(pos, parent)
			h.heapifyUp(parent)
		}
	}
}

func (h *Heap) heapifyDown(pos int) {
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

func init() {
	// do some initialization here
}

// end of file

package heap

type DHeap struct {
	degree int
	heap   []int
}

func NewDHeap(degree int) *DHeap {
	return &DHeap{
		degree: degree,
		heap:   []int{},
	}
}

func (h *DHeap) Insert(value int) {
	h.heap = append(h.heap, value)
	h.bubbleUp(len(h.heap) - 1)
}

func (h *DHeap) bubbleUp(i int) {
	parent := h.parent(i)
	for i > 0 && h.heap[parent] < h.heap[i] {
		h.heap[i], h.heap[parent] = h.heap[parent], h.heap[i]
		i = parent
		parent = h.parent(i)
	}
}

func (h *DHeap) parent(i int) int {
	return (i - 1) / h.degree
}

func (h *DHeap) ExtractMax() (int, bool) {
	if len(h.heap) == 0 {
		return 0, false
	}

	max := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]

	h.bubbleDown(0)

	return max, true
}

func (h *DHeap) bubbleDown(i int) {
	max := i
	children := h.childrenIndices(i)

	for _, child := range children {
		if child < len(h.heap) && h.heap[child] > h.heap[max] {
			max = child
		}
	}

	if max != i {
		h.heap[i], h.heap[max] = h.heap[max], h.heap[i]
		h.bubbleDown(max)
	}
}

func (h *DHeap) childrenIndices(i int) []int {
	children := make([]int, h.degree)
	for j := 0; j < h.degree; j++ {
		children[j] = h.degree*i + j + 1
	}
	return children
}

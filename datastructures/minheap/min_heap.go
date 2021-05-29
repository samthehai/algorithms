package minheap

type MinHeap struct {
	array []int
}

func NewMinHeap() MinHeap {
	return MinHeap{
		array: make([]int, 0),
	}
}

func (h *MinHeap) Size() int {
	return len(h.array)
}

func (h *MinHeap) PeekMin() *int {
	if h.Size() == 0 {
		return nil
	}

	min := h.array[0]

	return &min
}

func (h *MinHeap) ExtractMin() *int {
	minVal := h.PeekMin()

	if minVal == nil {
		return nil
	}

	n := h.Size()
	h.array[0] = h.array[n]
	h.array = h.array[:n-1]
	h.bubbleDown(0)

	return minVal
}

func (h *MinHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.bubbleUp(h.Size() - 1)
}

func (h *MinHeap) bubbleDown(index int) {
	indexSmallerChild := h.findSmallerChild(index)

	if indexSmallerChild >= h.Size() {
		return
	}

	if h.array[indexSmallerChild] < h.array[index] {
		h.array[indexSmallerChild], h.array[index] = h.array[index], h.array[indexSmallerChild]
		h.bubbleDown(indexSmallerChild)
	}
}

func (h *MinHeap) findSmallerChild(index int) (childIndex int) {
	size := h.Size()
	if index >= size {
		return -1
	}

	leftChildIndex, rightChildIndex := index*2+1, index*2+2
	if rightChildIndex >= size {
		if leftChildIndex >= size {
			return -1
		} else {
			return leftChildIndex
		}
	} else {
		if leftChildIndex < size && h.array[leftChildIndex] <= h.array[rightChildIndex] {
			return leftChildIndex
		} else {
			return rightChildIndex
		}
	}
}

func (h *MinHeap) bubbleUp(index int) {
	if index <= 0 {
		return
	}

	parentIndex := (index - 1) / 2
	if h.array[index] < h.array[parentIndex] {
		h.array[index], h.array[parentIndex] = h.array[parentIndex], h.array[index]
		h.bubbleUp(parentIndex)
	}
}

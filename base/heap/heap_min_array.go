package heap

import "fmt"

type minHeap struct {
	Capacity int
	Array    []int
}

func NewEmptyMinHeap() Heap {
	return &minHeap{
		Capacity: 0,
		Array:    make([]int, 0),
	}
}

func (h *minHeap) leftIndex(i int) int {
	return 2*i + 1
}

func (h *minHeap) rightIndex(i int) int {
	return 2*i + 2
}

func (h *minHeap) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *minHeap) AddElement(e int) {
	h.Array = append(h.Array, e)
	h.Capacity++
	j := h.Capacity - 1
	for j > 0 {
		if h.Array[j] < h.Array[h.parentIndex(j)] {
			h.Array[j], h.Array[h.parentIndex(j)] = h.Array[h.parentIndex(j)], h.Array[j]
		}
		j = h.parentIndex(j)
	}
}

func (h *minHeap) minIndexInTripe(i int) int {
	min := i
	l := h.leftIndex(i)
	r := h.rightIndex(i)
	if h.Array[l] < h.Array[min] {
		min = l
	}
	if h.Array[r] < h.Array[min] {
		min = r
	}
	return min
}

func (h *minHeap) RemoveElement() error {
	if h.Capacity <= 0 {
		return HeapEmptyError
	}
	h.Array[0] = h.Array[h.Capacity-1]
	h.Capacity--
	for i := 0; h.leftIndex(i) < h.Capacity; {
		min := h.minIndexInTripe(i)
		if i != min {
			h.Array[i], h.Array[min] = h.Array[min], h.Array[i]
			i = min
		}
	}
	return nil
}

func (h *minHeap) TopElement() (int, error) {
	if h.Capacity > 0 {
		return h.Array[0], nil
	}
	return 0, HeapEmptyError
}

func (h *minHeap) Print() {
	for i := 0; i < h.Capacity; i++ {
		fmt.Printf("%d:%d ", i, h.Array[i])
	}
	fmt.Println()
}

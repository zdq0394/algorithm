package heap

import "fmt"

type maxHeap struct {
	Capacity int
	Array    []int
}

func NewEmptyMaxHeap() Heap {
	return &maxHeap{
		Capacity: 0,
		Array:    make([]int, 0),
	}
}

func (h *maxHeap) AddElement(e int) {
	h.Array = append(h.Array, e)
	h.Capacity++
	j := h.Capacity - 1
	for j > 0 {
		if h.Array[j] > h.Array[h.parentIndex(j)] {
			h.Array[j], h.Array[h.parentIndex(j)] = h.Array[h.parentIndex(j)], h.Array[j]
		}
		j = h.parentIndex(j)
	}
}

func (h *maxHeap) RemoveElement() error {
	if h.Capacity <= 0 {
		return HeapEmptyError
	}
	h.Array[0] = h.Array[h.Capacity-1]
	h.Capacity--
	for i := 0; h.leftIndex(i) < h.Capacity; {
		max := h.maxIndexInTripe(i)
		if i != max {
			h.Array[i], h.Array[max] = h.Array[max], h.Array[i]
			i = max
		} else {
			break
		}
	}
	return nil
}

func (h *maxHeap) TopElement() (int, error) {
	if h.Capacity > 0 {
		return h.Array[0], nil
	}
	return 0, HeapEmptyError
}

func (h *maxHeap) Print() {
	for i := 0; i < h.Capacity; i++ {
		fmt.Printf("%d:%d ", i, h.Array[i])
	}
	fmt.Println()
}

func (h *maxHeap) leftIndex(i int) int {
	return 2*i + 1
}

func (h *maxHeap) rightIndex(i int) int {
	return 2*i + 2
}

func (h *maxHeap) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *maxHeap) maxIndexInTripe(i int) int {
	ret := i
	l := h.leftIndex(i)
	r := h.rightIndex(i)
	if l < h.Capacity && h.Array[l] > h.Array[ret] {
		ret = l
	}
	if r < h.Capacity && h.Array[r] > h.Array[ret] {
		ret = r
	}
	return ret
}

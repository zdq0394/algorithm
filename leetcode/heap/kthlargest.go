package heap

type KthLargest struct {
	Capacity int // size
	Array    []int
}

func Constructor(k int, nums []int) KthLargest {
	h := KthLargest{
		Capacity: k,
		Array:    make([]int, 0),
	}
	for _, v := range nums {
		h.Add(v)
	}
	return h
}

func (h *KthLargest) Add(e int) int {
	if len(h.Array) < h.Capacity {
		h.Array = append(h.Array, e)
		j := len(h.Array) - 1
		for j > 0 {
			if h.Array[j] < h.Array[h.parentIndex(j)] {
				h.Array[j], h.Array[h.parentIndex(j)] = h.Array[h.parentIndex(j)], h.Array[j]
				j = h.parentIndex(j)
			} else {
				break
			}
		}
	} else if e > h.Array[0] {
		h.Array[0] = e
		i := 0
		for h.leftIndex(i) < len(h.Array) {
			min := h.minIndexInTripe(i)
			if i != min {
				h.Array[i], h.Array[min] = h.Array[min], h.Array[i]
				i = min
			} else {
				break
			}
		}
	}
	return h.Array[0]
}

func (h *KthLargest) leftIndex(i int) int {
	return 2*i + 1
}

func (h *KthLargest) rightIndex(i int) int {
	return 2*i + 2
}

func (h *KthLargest) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *KthLargest) minIndexInTripe(i int) int {
	ret := i
	l := h.leftIndex(i)
	r := h.rightIndex(i)
	if l < len(h.Array) && h.Array[l] < h.Array[ret] {
		ret = l
	}
	if r < len(h.Array) && h.Array[r] < h.Array[ret] {
		ret = r
	}
	return ret
}

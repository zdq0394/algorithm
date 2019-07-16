package p10

func lastStoneWeight(stones []int) int {
	h := NewMaxHeap(len(stones))
	for i := 0; i < len(stones); i++ {
		h.Add(stones[i])
	}
	for {
		if h.size == 0 {
			return 0
		}
		if h.size == 1 {
			return h.data[0]
		}
		s1 := h.Remove()
		s2 := h.Remove()
		if s1 != s2 {
			h.Add(s1 - s2)
		}
	}
}

type maxheap struct {
	data []int
	size int
}

func NewMaxHeap(num int) *maxheap {
	return &maxheap{
		data: make([]int, num),
		size: 0,
	}
}

func (h *maxheap) Add(n int) {
	h.data[h.size] = n
	h.size++
	curIndex := h.size - 1
	for curIndex > 0 {
		if h.data[curIndex] > h.data[(curIndex-1)/2] {
			h.data[curIndex], h.data[(curIndex-1)/2] = h.data[(curIndex-1)/2], h.data[curIndex]
			curIndex = (curIndex - 1) / 2
		} else {
			break
		}
	}
}

func (h *maxheap) Remove() int {
	ret := h.data[0]
	h.size--
	h.data[0] = h.data[h.size]
	curIndex := 0
	for 2*curIndex+1 <= h.size-1 {
		adjustIndex := 2*curIndex + 1
		if 2*curIndex+2 <= h.size-1 {
			if h.data[curIndex] >= h.data[2*curIndex+1] && h.data[curIndex] >= h.data[2*curIndex+2] {
				break
			} else {
				if h.data[2*curIndex+1] < h.data[2*curIndex+2] {
					adjustIndex = 2*curIndex + 2
				}
			}
		} else {
			if h.data[curIndex] > h.data[2*curIndex+1] {
				break
			}
		}
		h.data[curIndex], h.data[adjustIndex] = h.data[adjustIndex], h.data[curIndex]
		curIndex = adjustIndex
	}
	return ret
}

package heap

func findKthLargest(nums []int, k int) int {
	heap := []int{0} // index 0 will be ignored.
	for _, v := range nums {
		if len(heap) < k+1 {
			heap = append(heap, v)
			j := len(heap) - 1
			for j > 1 {
				p := j / 2
				if heap[j] < heap[p] {
					heap[j], heap[p] = heap[p], heap[j]
					j = p
				} else {
					break
				}
			}
		} else {
			if v > heap[1] {
				heap[1] = v
				j := 1
				for 2*j < len(heap) {
					min := findMinInTriple(heap, j)
					if j != min {
						heap[j], heap[min] = heap[min], heap[j]
						j = min
					} else {
						break
					}
				}
			}
		}
	}
	return heap[1]
}

func findMinInTriple(heap []int, i int) int {
	ret := i
	if 2*i < len(heap) && heap[2*i] < heap[ret] {
		ret = 2 * i
	}
	if 2*i+1 < len(heap) && heap[2*i+1] < heap[ret] {
		ret = 2*i + 1
	}
	return ret
}

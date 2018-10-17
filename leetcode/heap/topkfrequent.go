package heap

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, k := range nums {
		m[k]++
	}
	values := []pair{}
	for k, v := range m {
		values = append(values, pair{k, v})
	}
	topKPairs := findTopKLargest1(values, k)
	topK := []int{}
	for _, t := range topKPairs {
		topK = append(topK, t.e)
	}
	return topK
}

type pair struct {
	e   int
	fre int
}

func findTopKLargest1(nums []pair, k int) []pair {
	heap := []pair{pair{}} // index 0 will be ignored.
	for _, v := range nums {
		if len(heap) < k+1 {
			heap = append(heap, v)
			j := len(heap) - 1
			for j > 1 {
				p := j / 2
				if heap[j].fre < heap[p].fre {
					heap[j], heap[p] = heap[p], heap[j]
					j = p
				} else {
					break
				}
			}
		} else {
			if v.fre > heap[1].fre {
				heap[1] = v
				j := 1
				for 2*j < len(heap) {
					min := findMinInTriple1(heap, j)
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
	return heap[1:]
}

func findMinInTriple1(heap []pair, i int) int {
	ret := i
	if 2*i < len(heap) && heap[2*i].fre < heap[ret].fre {
		ret = 2 * i
	}
	if 2*i+1 < len(heap) && heap[2*i+1].fre < heap[ret].fre {
		ret = 2*i + 1
	}
	return ret
}

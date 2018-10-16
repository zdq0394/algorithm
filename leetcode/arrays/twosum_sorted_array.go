package arrays

func twoSumSorted1(numbers []int, target int) []int {
	m := make(map[int]int)
	for i, v := range numbers {
		vv := target - v
		if j, ok := m[vv]; ok && j != i {
			return []int{j, i}
		}
		m[v] = i
	}
	return []int{}
}

func twoSumSorted2(numbers []int, target int) []int {
	min := 0
	max := len(numbers) - 1
	for min < max {
		val := numbers[min] + numbers[max]
		if val < target {
			min++
		} else if val > target {
			max--
		} else {
			return []int{min, max}
		}
	}
	return []int{}
}

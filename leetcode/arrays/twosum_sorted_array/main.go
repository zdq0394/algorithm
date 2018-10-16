package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for i, v := range numbers {
		vv := target - v
		if j, ok := m[vv]; ok && j != i {
			return []int{j + 1, i + 1}
		}
		m[v] = i
	}
	return []int{}
}

func twoSum2(numbers []int, target int) []int {
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

func main() {
	target := 11
	nums := []int{5, 8, 2, 6, 0, 7}
	result := twoSum(nums, target)
	fmt.Println(result)
}

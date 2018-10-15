package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	if nums == nil {
		return -1
	}
	h := map[int]bool{}
	for _, v := range nums {
		h[v] = true
	}
	for i := 0; i <= len(nums); i++ {
		if _, ok := h[i]; !ok {
			return i
		}
	}
	return -1
}

func missingNumber1(nums []int) int {
	if nums == nil {
		return -1
	}
	n := len(nums)
	if n == 0 {
		return 0
	}
	total := n * (n + 1) / 2
	for _, v := range nums {
		total -= v
	}
	return total
}

func main() {
	nums := []int{0, 1, 2, 3}
	fmt.Println(missingNumber(nums))
	fmt.Println(missingNumber1(nums))
}

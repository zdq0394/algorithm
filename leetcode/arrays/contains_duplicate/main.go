package main

import "fmt"

func containsDuplicate(nums []int) bool {
	if nums == nil || len(nums) < 2 {
		return false
	}
	h := map[int]bool{}
	for _, v := range nums {
		if _, ok := h[v]; ok {
			return true
		} else {
			h[v] = true
		}
	}
	return false
}

func containsNearbyDuplicate(nums []int, k int) bool {
	if nums == nil || len(nums) < 2 {
		return false
	}
	h := map[int]int{}
	for i, v := range nums {
		if j, ok := h[v]; ok {
			if i-j <= k {
				return true
			}
			h[v] = i
		} else {
			h[v] = i
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 5, 4, 3}
	fmt.Println(containsDuplicate(nums))
	fmt.Println(containsNearbyDuplicate(nums, 4))
}

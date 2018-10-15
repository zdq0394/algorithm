package main

import "fmt"

func majorityElement(nums []int) int {
	h := make(map[int]int, 0)
	l := len(nums)
	for _, v := range nums {
		if n, ok := h[v]; ok {
			h[v] = n + 1
			if h[v] > l/2 {
				return v
			}
		} else {
			h[v] = 1
		}
	}
	return 0
}

func majorityElement2(nums []int) int {
	l := len(nums)
	r := nums[0]
	c := 1
	for i := 1; i < l; i++ {
		if nums[i] == r {
			c++
			if c > l/2 {
				return r
			}
		} else {
			if c == 1 {
				r = nums[i]
			} else {
				c--
			}
		}

	}
	return r
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
	fmt.Println(majorityElement2(nums))
}

package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	t := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if t > max {
				max = t
			}
			t = 0
		} else {
			t++
		}
	}
	if t > max {
		max = t
	}
	return max
}

func main() {
	nums := []int{1, 2, 1, 1, 1, 4, 5, 1, 2, 1, 1, 1, 1, 0}
	fmt.Println(findMaxConsecutiveOnes(nums))
}

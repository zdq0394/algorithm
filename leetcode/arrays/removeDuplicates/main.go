package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if nums == nil {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	lastIndex := 0
	cursorIndex := 1
	orignalDiff := 1
	for {
		if nums[cursorIndex] == nums[lastIndex] {
			cursorIndex++
		} else {
			diff := cursorIndex - lastIndex + orignalDiff
			if diff > 1 {
				nums[lastIndex+1] = nums[cursorIndex]
			}
			lastIndex++
			cursorIndex++
		}
		if cursorIndex == len(nums) {
			break
		}
	}
	fmt.Println(nums)
	return lastIndex + 1
}

func main() {
	nums := []int{1, 1, 2, 3, 3, 3, 3, 4, 4, 5, 6, 7, 8, 8}
	fmt.Println(removeDuplicates(nums))
}

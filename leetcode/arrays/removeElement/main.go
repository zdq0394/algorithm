package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	if nums == nil {
		return 0
	}
	if len(nums) == 0 {
		return 0
	}
	lastIndex := -1
	cursorIndex := 0
	delta := 1
	for {
		if nums[cursorIndex] == val {
			cursorIndex++
		} else {
			diff := cursorIndex - lastIndex + delta
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
	fmt.Println(removeElement(nums, 3))
}

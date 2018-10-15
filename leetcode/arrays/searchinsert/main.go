package main

import (
	"fmt"
	"time"
)

func searchInsert1(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	mid := len(nums) / 2
	switch {
	case target == nums[mid]:
		return mid
	case target < nums[mid]:
		if mid == 0 {
			return mid
		}
		if target > nums[mid-1] {
			return mid
		}
		if target <= nums[mid-1] {
			return searchInsert1(nums[:mid], target)
		}
	case target > nums[mid]:
		if mid == len(nums)-1 {
			return mid + 1
		}
		if target < nums[mid+1] {
			return mid + 1
		}
		if target >= nums[mid+1] {
			return mid + 1 + searchInsert1(nums[mid+1:], target)
		}
	}
	return -1
}

func searchInsert2(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
		if v > target {
			return i
		}
	}
	return len(nums)
}

func main() {
	target := 128
	var nums [100000000]int
	time1 := time.Now()
	fmt.Println(searchInsert1(nums[:], target))
	time2 := time.Now()
	fmt.Println(searchInsert2(nums[:], target))
	time3 := time.Now()

	delta2 := time3.Sub(time2)
	delta1 := time2.Sub(time1)
	fmt.Println(delta1)
	fmt.Println(delta2)
}

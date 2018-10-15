package main

import "fmt"

func moveZeroes(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}
	move := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i-move] = nums[i]
		} else {
			move++
		}
	}
	for j := len(nums) - move; j <= len(nums)-1; j++ {
		nums[j] = 0
	}
}
func main() {
	nums := []int{0, 1, 2, 3}
	moveZeroes(nums)
	fmt.Println(nums)
}

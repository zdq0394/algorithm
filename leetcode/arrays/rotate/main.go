package main

import "fmt"

func rotate(nums []int, k int) {
	l := len(nums)
	for k > 0 {
		t := nums[l-1]
		for i := l - 1; i >= 1; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = t
		fmt.Println(nums)
		k--
	}
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(nums)
	k := 7
	rotate(nums, k)
	fmt.Printf("Move %d towards right:\n%v", k, nums)
}

package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	y := 0
	orignalX := x
	for {
		y = y*10 + x%10
		x /= 10
		if x == 0 {
			break
		}
	}
	if orignalX == y {
		fmt.Println("Yes")
		return true
	}
	fmt.Println("No")
	return false
}

func main() {
	isPalindrome(121383121)
}

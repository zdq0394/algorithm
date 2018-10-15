package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	plus := 1
	for j := len(digits) - 1; j >= 0; j-- {
		digits[j] += plus
		plus = digits[j] / 10
		digits[j] = digits[j] % 10
		if plus == 0 {
			break
		}
	}
	if plus == 1 {
		t := []int{1}
		digits = append(t, digits...)
	}

	return digits
}

func main() {
	i := []int{9, 9, 9, 9}
	fmt.Print(plusOne(i))
}

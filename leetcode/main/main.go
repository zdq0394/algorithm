package main

import (
	"fmt"

	"github.com/zdq0394/algorithm/leetcode/arrays"
)

func main() {
	a := []int{0, 0, 9, 0, 0, 0, 2, 1}
	arrays.DuplicateZeros(a)
	fmt.Println(a)
}

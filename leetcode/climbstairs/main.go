package main

import (
	"fmt"
	"time"
)

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-2) + climbStairs(n-1)
}

func climbStairs1(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	result := make([]int, n)
	result[0] = 1
	result[1] = 2

	for i := 2; i < n; i++ {
		result[i] = result[i-2] + result[i-1]
	}

	return result[n-1]
}

func main() {
	n := 42
	t1 := time.Now()
	fmt.Println(climbStairs(n))
	t2 := time.Now()
	fmt.Println(climbStairs1(n))
	t3 := time.Now()
	fmt.Println("Recur:", t2.Sub(t1))
	fmt.Println("Loop:", t3.Sub(t2))
}

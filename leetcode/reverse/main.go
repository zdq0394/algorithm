package main

import (
	"fmt"
	"math"
)

func reverse1(x int) int {
	if x < 10 && x > -10 {
		return x
	}

	sign := 1
	if x <= -10 {
		x = -x
		sign = -1
	}

	var numbers [10]int
	var i int
	for ; i < 10; i++ {
		if x/10 == 0 {
			numbers[i] = x
			break
		}
		remainder := x / 10
		last := x - x/10*10
		x = x / 10
		numbers[i] = last
		x = remainder
	}

	validNums := numbers[:i+1]

	var tempNum int64
	for j := 0; j < len(validNums); j++ {
		if validNums[j] == 0 {
			continue
		}
		tempNum += int64(math.Pow10(len(validNums)-1-j)) * int64(validNums[j])
	}
	if tempNum > int64(math.Pow(2, 31)-1) {
		return 0
	}
	var ret int64
	ret = tempNum * int64(sign)
	if ret > math.MaxInt32 || ret < math.MinInt32 {
		return 0
	} else {
		return int(ret)
	}
}

func reverse2(x int) int {
	y := 0
	for {
		y = y*10 + (x % 10)
		if x/10 == 0 {
			break
		}
		x = x / 10
	}
	return y
}

func main() {
	ret := reverse2(1456789)
	fmt.Println(ret)
}

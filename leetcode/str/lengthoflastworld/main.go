package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	lastValidIndex := len(s) - 1
	for j := len(s) - 1; j >= 0; j-- {
		if s[j] == ' ' {
			lastValidIndex--
		} else {
			break
		}
	}
	var k int
	for k = lastValidIndex; k >= 0; k-- {
		if s[k:k+1] == " " {
			break
		}
	}
	return lastValidIndex - k
}

func main() {
	fmt.Println(lengthOfLastWord("         "))
}

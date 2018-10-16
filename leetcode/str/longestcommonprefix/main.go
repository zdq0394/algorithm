package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	l := len(strs[0])

	for _, str := range strs {
		if len(str) < l {
			l = len(str)
		}
	}

	maxIndex := 0
	done := true
	for i := 0; done && i < l; i++ {
		for _, str := range strs[1:] {
			if strs[0][i] != str[i] {
				done = false
				break
			}
		}
		maxIndex = i
	}

	if done {
		maxIndex += 1
	}

	if maxIndex == 0 {
		return ""
	}
	if l != 0 {
		return strs[0][:maxIndex]
	}
	return ""

}

func main() {
	s := []string{"", ""}
	fmt.Println(longestCommonPrefix(s))
}

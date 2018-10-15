package main

import "fmt"

func strStr(haystack string, needle string) int {
	l1 := len(haystack)
	l2 := len(needle)

	for i := 0; i+l2 <= l1; i++ {
		if string(haystack[i:i+l2]) == needle {
			return i
		}
	}
	return -1
}
func main() {
	s1 := "aaabbbcdefff"
	s2 := "bbcd"
	fmt.Println(strStr(s1, s2))
}

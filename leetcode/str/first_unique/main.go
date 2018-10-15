package main

import "fmt"

func firstUniqChar(s string) int {
	var ss [26]int
	for i := 0; i < len(s); i++ {
		ss[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if ss[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("z"))
}

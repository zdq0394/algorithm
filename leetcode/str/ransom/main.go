package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	h := [26]int{}
	for i := 0; i < len(magazine); i++ {
		h[magazine[i]-'a']++

	}

	for j := 0; j < len(ransomNote); j++ {
		h[ransomNote[j]-'a']--
		if h[ransomNote[j]-'a'] < 0 {
			return false
		}
	}

	return true
}

func main() {
	s1 := "aa"
	s2 := "aab"
	fmt.Println(canConstruct(s1, s2))
}

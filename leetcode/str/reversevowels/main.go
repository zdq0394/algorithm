package main

import "fmt"

func validVowels(c byte) bool {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	_, ok := vowels[c]
	return ok
}
func reverseVowels(s string) string {
	ss := []byte(s)
	i, j := 0, len(ss)-1
	for i <= j {
		for !validVowels(ss[i]) {
			i++
			if i >= j {
				break
			}
		}
		for !validVowels(ss[j]) {
			j--
			if j <= i {
				break
			}
		}
		if i < j {
			ss[i], ss[j] = ss[j], ss[i]
		}
		i++
		j--
	}
	return string(ss)
}
func main() {
	fmt.Println(reverseVowels("leetcode"))
}

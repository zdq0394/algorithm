package main

import "fmt"

func isValidChar(c byte) (byte, bool) {
	if c >= '0' && c <= '9' {
		return c, true
	}
	if c >= 'a' && c <= 'z' {
		return c, true
	}
	if c >= 'A' && c <= 'Z' {
		return c + byte(32), true
	}
	return 0, false
}
func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		ic, iOK := isValidChar(s[i])
		jc, jOK := isValidChar(s[j])
		if iOK && jOK {
			if ic != jc {
				return false
			}
			i++
			j--
		} else {
			if !iOK {
				i++
			}
			if !jOK {
				j--
			}
		}
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
	fmt.Printf("A:%d;a%d", 'A', 'a')
}

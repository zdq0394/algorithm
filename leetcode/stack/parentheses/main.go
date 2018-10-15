package main

import "fmt"

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) == 1 {
		return false
	}
	l := len(s)
	b := make([]byte, l)
	next := 0
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			b[next] = byte(c)
			next++
		case ')':
			if next == 0 {
				return false
			}
			if b[next-1] == '(' {
				next = next - 1
			} else {
				return false
			}
		case ']':
			if next == 0 {
				return false
			}
			if b[next-1] == '[' {
				next = next - 1
			} else {
				return false
			}
		case '}':
			if next == 0 {
				return false
			}
			if b[next-1] == '{' {
				next = next - 1
			} else {
				return false
			}
		}
	}
	return next == 0
}

func main() {
	fmt.Println(isValid("){"))
}

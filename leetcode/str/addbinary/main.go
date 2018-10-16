package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	s := ""
	var plus byte
	for i >= 0 || j >= 0 {
		t := plus
		if i >= 0 {
			t = t + a[i] - byte(48)
		}
		if j >= 0 {
			t = t + b[j] - byte(48)
		}
		r := t % 2
		plus = t / 2
		s = fmt.Sprintf("%d%s", r, s)
		i--
		j--
	}

	if plus > 0 {
		s = fmt.Sprintf("%d%s", 1, s)
	}
	return s
}

func main() {
	fmt.Println(addBinary("1101", "101"))
}

package main

import (
	"fmt"
)

func reverseString(s string) string {
	bs := []byte(s)
	i, j := 0, len(bs)-1
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return string(bs)
}

func main() {
	fmt.Println(reverseString("Hello"))
}

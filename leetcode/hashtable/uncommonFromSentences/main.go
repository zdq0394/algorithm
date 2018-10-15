package main

import "fmt"

func readString(a string, h map[string]int) {
	begin := 0
	for i := begin + 1; i < len(a); i++ {
		if a[i] == ' ' || i == len(a)-1 {
			word := a[begin:i]
			if i == len(a)-1 {
				word = a[begin:len(a)]
			}
			if word != " " {
				h[word]++
			}
			begin = i + 1
		}
	}
}
func uncommonFromSentences(a string, b string) []string {
	result := make([]string, 0)
	h := make(map[string]int)
	readString(a, h)
	readString(b, h)
	for k, v := range h {
		if v == 1 {
			result = append(result, k)
		}
	}
	return result
}
func main() {
	a := "this apple is sweet"
	b := "this apple is sour"
	fmt.Println(uncommonFromSentences(a, b))
}

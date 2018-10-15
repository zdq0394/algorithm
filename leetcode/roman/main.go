package main

import "fmt"

func romanToInt1(roman string) int {
	m := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}
	if len(roman) == 1 {
		return m[roman]
	}

	value := 0

	for i := range roman {
		if i == len(roman)-1 {
			value += m[string(roman[i])]
		} else {
			curC := string(roman[i])
			nextC := string(roman[i+1])
			if m[curC] < m[nextC] {
				value -= m[curC]
			} else {
				value += m[curC]
			}
		}
	}
	return value
}

func main() {
	fmt.Println(romanToInt1("MCMXCIV"))
}

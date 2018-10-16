package main

import "fmt"

func countAndSay2(n int) string {
	// it is the look-and-say sequence 康威
	if n == 1 {
		return "1"
	} else {
		toCountString := countAndSay(n - 1)
		var curChr, curCnt rune
		m := make([]rune, 0)
		for _, s := range toCountString {
			if curChr != s {
				if curChr != 0 {
					m = append(m, curCnt, curChr)
				}
				curChr = s
				curCnt = 49
			} else {
				curCnt++
			}
		}

		m = append(m, curCnt, curChr)
		return string(m)
	}
}

func countAndSay(n int) string {
	m := make(map[int]string, 0)
	m[1] = "1"
	m[2] = "11"
	for i := 3; i <= n; i++ {
		m[i] = readANumber(m[i-1])
	}
	return m[n]
}

func readANumber(num string) string {
	if len(num) == 0 {
		return ""
	}
	if len(num) == 1 {
		return "1" + num
	}
	c := num[0]
	count := 1
	result := ""
	for i := 1; i < len(num); i++ {
		if num[i] == c {
			count++
			if i == len(num)-1 {
				result = fmt.Sprintf("%s%d%c", result, count, c)
			}
		} else {
			result = fmt.Sprintf("%s%d%c", result, count, c)
			c = num[i]
			count = 1
			if i == len(num)-1 {
				result = fmt.Sprintf("%s%d%c", result, count, c)
			}
		}
	}
	return result
}

func main() {
	n := 8
	fmt.Println(countAndSay(n))
	fmt.Println(countAndSay2(n))
}

package arrays

import "fmt"

func findShortestSubArray(nums []int) int {
	h := map[int]Degree{}
	for i := 0; i < len(nums); i++ {
		if d, ok := h[nums[i]]; ok {
			d.value++
			d.end = i
			h[nums[i]] = d
		} else {
			h[nums[i]] = Degree{
				value: 1,
				start: i,
				end:   i,
			}
		}
	}
	maxDegree := 0
	maxDegreeMinLength := len(nums)
	for _, d := range h {
		if d.value > maxDegree {
			maxDegree = d.value
			maxDegreeMinLength = d.end - d.start + 1
		} else if d.value == maxDegree {
			if d.end-d.start+1 < maxDegreeMinLength {
				maxDegreeMinLength = d.end - d.start + 1
			}
		}
		fmt.Println("maxDegree:", maxDegree)
		fmt.Println("maxDegreeLength:", maxDegreeMinLength)
	}
	return maxDegreeMinLength
}

type Degree struct {
	value int
	start int
	end   int
}

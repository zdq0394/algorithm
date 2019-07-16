package leetcode

func numPairsDivisibleBy60(time []int) int {
	if time == nil || len(time) == 0 {
		return 0
	}
	h := map[int]int{}
	t := 0
	for i := 0; i < len(time); i++ {
		time[i] = time[i] % 60
		if time[i] == 0 {
			if v, ok := h[time[i]]; ok {
				t += v
			}
		} else {
			if v, ok := h[60-time[i]]; ok {
				t += v
			}
		}
		h[time[i]]++
	}
	return t
}

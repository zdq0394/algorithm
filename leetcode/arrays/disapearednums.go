package arrays

func findDisappearedNumbers(nums []int) []int {
	retNums := make([]int, len(nums)+1)
	for _, n := range nums {
		retNums[n] = 1
	}
	var j int
	for i, r := range retNums[1:] {
		if r == 0 {
			retNums[j] = i + 1
			j++
		}
	}
	return retNums[:j]
}

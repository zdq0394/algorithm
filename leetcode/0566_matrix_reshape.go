package leetcode

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if nums == nil || len(nums) == 0 || len(nums[0]) == 0 {
		return nums
	}
	m := len(nums)
	n := len(nums[0])
	if m*n != r*c {
		return nums
	}
	newNums := make([][]int, r)
	for i := 0; i < r; i++ {
		newNums[i] = make([]int, c)
		for j := 0; j < c; j++ {
			t := i*c + j
			newNums[i][j] = nums[t/n][t%n]
		}
	}
	return newNums
}

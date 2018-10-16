package arrays

func twoSum1(nums []int, target int) []int {
	for i, n := range nums {
		left := target - n
		for j := len(nums) - 1; j > i; j-- {
			if left == nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		vv := target - v
		if j, ok := m[vv]; ok && j != i {
			return []int{j, i}
		}
		m[v] = i
	}
	return []int{}
}

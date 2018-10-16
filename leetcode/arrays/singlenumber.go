package arrays

func singleNumber(nums []int) int {
	var e int
	for _, v := range nums {
		e = e ^ v
	}
	return e
}

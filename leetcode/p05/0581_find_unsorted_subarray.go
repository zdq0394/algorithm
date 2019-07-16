package p05

func findUnsortedSubarray(nums []int) int {
	if nums == nil || len(nums) == 0 || len(nums) == 1 {
		return 0
	}
	a := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		a[i] = nums[i]
	}
	qs(a, 0, len(a)-1)
	start := 0
	var i int
	for i = 0; i < len(a); i++ {
		if a[i] != nums[i] {
			start = i
			break
		}
	}
	if i == len(a) {
		return 0
	}
	end := 0
	var j int
	for j = len(a) - 1; j >= 0; j-- {
		if a[j] != nums[j] {
			end = j
			break
		}
	}
	if j == -1 {
		return 0
	}
	return end - start + 1
}

func qs(a []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	p := a[left]
	for i < j {
		for i < j && a[j] >= p {
			j--
		}
		if i < j {
			a[i] = a[j]
		}
		for i < j && a[i] <= p {
			i++
		}
		if i < j {
			a[j] = a[i]
		}
	}
	a[i] = p
	qs(a, left, i-1)
	qs(a, i+1, right)
}

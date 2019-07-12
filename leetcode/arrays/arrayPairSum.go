package arrays

func ArrayPairSum(nums []int) int {
	quickSort(nums, 0, len(nums)-1)
	sum := 0
	for i := 0; i < len(nums)-1; {
		sum += nums[i]
		i += 2
	}
	return sum
}

func quickSort(a []int, l, r int) {
	if l >= r {
		return
	}
	i, j := l, r
	p := a[i]
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
		if j > i {
			a[j] = a[i]
		}
	}
	a[i] = p
	quickSort(a, l, i-1)
	quickSort(a, i+1, r)
}

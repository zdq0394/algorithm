package sort

func quickSort(ori []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	p := ori[i]
	for i < j {
		for i < j && ori[j] >= p {
			j--
		}
		if i < j {
			ori[i] = ori[j]
		}
		for i < j && ori[i] <= p {
			i++
		}
		if i < j {
			ori[j] = ori[i]
		}
	}
	ori[i] = p
	quickSort(ori, left, i-1)
	quickSort(ori, i+1, right)
}

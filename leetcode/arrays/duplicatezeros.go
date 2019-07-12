package arrays

func DuplicateZeros(arr []int) {
	k := len(arr) - 2
	for k >= 0 {
		if arr[k] == 0 {
			for i := len(arr) - 1; i > k; i-- {
				arr[i] = arr[i-1]
			}
		}
		k--
	}
}

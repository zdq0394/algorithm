package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	pos := m + n - 1
	m = m - 1
	n = n - 1
	for m >= 0 && n >= 0 {
		if nums1[m] >= nums2[n] {
			nums1[pos] = nums1[m]
			m--
		} else {
			nums1[pos] = nums2[n]
			n--
		}
		pos--
	}
	for ; n >= 0; n-- {
		nums1[n] = nums2[n]
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0}
	nums2 := []int{2, 4}
	merge(nums1, 3, nums2, 2)
	fmt.Println(nums1)
}

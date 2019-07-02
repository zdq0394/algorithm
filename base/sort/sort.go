package sort

var result = []int{0, 3, 3, 3, 4, 5, 6, 6, 8, 9}
var arr = []int{0, 3, 3, 3, 4, 5, 9, 6, 6, 8}

func equals(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

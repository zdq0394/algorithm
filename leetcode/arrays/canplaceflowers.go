package arrays

func canPlaceFlowers(flowerbed []int, n int) bool {
	pre := 0
	canPlace := 0
	a := flowerbed
	var i int
	for ; i < len(a)-1; i++ {
		if pre == 0 {
			if a[i] == 0 {
				if a[i+1] == 0 {
					canPlace++
					a[i] = 1
				}
				pre = a[i]
			} else {
				pre = a[i]
			}
		} else {
			pre = a[i]
		}
	}
	if pre == 0 && a[i] == 0 {
		canPlace++
	}
	return canPlace == n
}

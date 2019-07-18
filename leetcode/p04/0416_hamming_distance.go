package p04

func hammingDistance(x int, y int) int {
	x32 := int32(x)
	y32 := int32(y)
	z := x32 ^ y32
	c := 0
	if z < 0 {
		z = z & 0x7fffffff
		c++
	}
	for z != 0 {
		if z&1 == 1 {
			c++
		}
		z = z >> 1
	}
	return c
}

package p06

func judgeCircle(moves string) bool {
	x := 0
	y := 0
	for _, v := range moves {
		switch v {
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x++
		case 'R':
			x--
		}
	}
	return x == 0 && y == 0
}

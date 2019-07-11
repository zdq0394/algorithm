package arrays

func isOneBitCharacter(bits []int) bool {
	if bits == nil || len(bits) == 0 {
		return false
	}
	i := 0
	for i < len(bits)-1 {
		i += (bits[i] + 1)
	}
	return i == len(bits)-1 && bits[i] == 0
}

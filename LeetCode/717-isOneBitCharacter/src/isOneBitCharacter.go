package src

func isOneBitCharacter(bits []int) bool {
	ones := 0
	for i := len(bits) - 2; i >= 0 && 1 == bits[i]; i-- {
		ones++
	}

	if ones%2 != 0 {
		return false
	}
	return true
}

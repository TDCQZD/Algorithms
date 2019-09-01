package src

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if 9 != digits[i] {
			digits[i]++
			break
		} else {
			digits[i] = 0
		}
	}
	if 0 == digits[0] {
		newAns := make([]int, len(digits)+1)
		newAns[0] = 1
		return newAns
	}
	return digits
}
func plusOne1(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}
	}
	if 0 == digits[0] {
		newAns := make([]int, len(digits)+1)
		newAns[0] = 1
		return newAns
	}
	return digits
}

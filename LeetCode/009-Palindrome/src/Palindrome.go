package src

func isPalindrome(x int) bool {
	if 0 > x || (0 == x%10 && 0 != x) { // 负数不是回文的
		return false
	}

	var revertedNumber int
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	return (x == revertedNumber) || (x == revertedNumber/10)
}

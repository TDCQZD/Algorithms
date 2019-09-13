package src

func reverseVowels(s string) string {
	res := []byte(s)
	match := map[byte]int{
		'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1,
		'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1,
	}
	l, r := 0, len(res)-1
	for l < r {
		a, b := match[res[l]], match[res[r]]
		if a == 1 && b == 1 {
			res[l], res[r] = res[r], res[l]
			l++
			r--
		} else if a == 0 {
			l++
		} else if b == 0 {
			r--
		} else {
			l++
			r--
		}
	}
	return string(res)

}
func reverseVowels2(s string) string {
	str := []rune(s)
	i, j := 0, len(str)-1
	for i < j {
		for ; i < j && !isVolume(str[i]); i++ {
		} // 从前往后找到第一个元音字母
		for ; i < j && !isVolume(str[j]); j-- {
		} // 从后往前找到第一个元音字母
		if i < j {
			str[i], str[j] = str[j], str[i]
			i, j = i+1, j-1
		}
	}

	return string(str)
}

func isVolume(c rune) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}

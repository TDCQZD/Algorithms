package src
func addStrings(num1 string, num2 string) string {
	str1 := []rune(num1)
	str2 := []rune(num2)
	ans := make([]rune, 0)
	var carry rune = 0
	for i, j := len(str1)-1, len(str2)-1; i >= 0 || j >= 0 || 1 == carry; i, j = i-1, j-1 {
		var x, y rune = 0, 0
		if i >= 0 {
			x = str1[i] - '0'
		}
		if j >= 0 {
			y = str2[j] - '0'
		}
		bit := (x+y+carry)%rune(10) + '0'
		carry = (x + y + carry) / rune(10)
		ans = append(ans, bit)
	}

	for i, j := 0, len(ans)-1; i < j; {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
	return string(ans)
}

func addStrings2(num1 string, num2 string) string {
	ans := make([]byte, 0)
	carry := byte(0)
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || byte(1) == carry; i, j = i-1, j-1 {
		var x, y byte
		if i >= 0 {
			x = num1[i] - '0'
		}
		if j >= 0 {
			y = num2[j] - '0'
		}
		bit := (x+y+carry)%10 + '0'
		carry = (x + y + carry) / 10
		ans = append(ans, bit)
	}

	for i, j := 0, len(ans)-1; i < j; {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
	return string(ans)
}

func addStrings3(s1, s2 string) string {
	// 确保 n1 <= n2
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	n1, n2 := len(s1), len(s2)
	a1, a2 := []byte(s1), []byte(s2)

	carry := byte(0)

	// buf 保存 []byte 格式的答案
	buf := make([]byte, n2+1)
	buf[0] = '1'

	i := 1
	for i <= n2 {
		// a1 和 a2 相加
		if i <= n1 {
			buf[n2+1-i] = a1[n1-i] - '0'
		}
		buf[n2+1-i] += a2[n2-i] + carry

		// 处理进位问题
		if buf[n2+1-i] > '9' {
			buf[n2+1-i] -= 10
			carry = byte(1)
		} else {
			carry = byte(0)
		}

		i++
	}

	if carry == 1 {
		return string(buf)
	}
	return string(buf[1:])
}
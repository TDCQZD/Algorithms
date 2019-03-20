package src

import (
	"fmt"
	"strings"
)

// 反转字符串单词
func reverseWords(s string) string {

	words := strings.Split(s, " ")
	var strs []string
	for i := len(words) - 1; i >= 0; i-- {
		fmt.Println(words[i])
		if "" != words[i] {
			strs = append(strs, words[i])
		}
	}
	return strings.Join(strs, " ")
}

// 反转字符串
func reverse(s string) string {
	if len(s) <= 0 {
		return ""
	}
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

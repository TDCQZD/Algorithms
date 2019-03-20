package src

import (
	"fmt"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	str := "Let's take LeetCode contest"
	fmt.Println(reverseWords(str))
}

package src

import (
	"fmt"
	"testing"
)

func Test_reverseStr(t *testing.T) {
	s := "abcdefg"
	k := 2

	fmt.Println(reverseStr(s, k))

	s1 := "a"
	k1 := 2
	fmt.Println(reverseStr(s1, k1))
}

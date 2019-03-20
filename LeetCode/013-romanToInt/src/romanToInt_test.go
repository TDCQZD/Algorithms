package src

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	s := "MCMXCIV"
	res := romanToInt(s)
	fmt.Println(res)
}

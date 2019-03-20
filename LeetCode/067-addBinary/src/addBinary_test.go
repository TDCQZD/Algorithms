package src

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	a := "1010"
	b := "1011"
	res := addBinary(a, b)
	fmt.Println(res)
	res = addBinary2(a, b)
	fmt.Println(res)
}

package src

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	arr := []int{0}
	res := plusOne(arr)
	fmt.Println(res)
	arr = []int{4, 3, 2, 1}
	res = plusOne(arr)
	fmt.Println(res)
}

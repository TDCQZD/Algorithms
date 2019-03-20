package src

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	arr := []int{1, 3, 5, 6}
	res := searchInsert(arr, 5)
	fmt.Println(res)
	res = searchInsert(arr, 2)
	fmt.Println(res)
	res = searchInsert(arr, 7)
	fmt.Println(res)
	res = searchInsert(arr, 0)
	fmt.Println(res)
}

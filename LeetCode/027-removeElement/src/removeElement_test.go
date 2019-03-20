package src

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{0,1,2,2,3,0,4,2}
	length, newNums := removeElement(nums, 2)
	fmt.Println(length)
	fmt.Println(newNums)
}

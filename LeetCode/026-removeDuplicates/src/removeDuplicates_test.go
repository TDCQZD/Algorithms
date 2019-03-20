package src

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{1, 1, 1, 2, 3, 3, 4}
	fmt.Println("length:", len(nums))
	len, newNums := removeDuplicates(nums)
	fmt.Println("new length:", len)
	fmt.Println("newNums:", newNums)

}

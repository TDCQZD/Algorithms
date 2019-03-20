package src

import (
	"fmt"
	"testing"
)

func Test_findErrorNums(t *testing.T) {
	nums := []int{4, 2, 1, 2}
	fmt.Println(findErrorNums(nums))
}

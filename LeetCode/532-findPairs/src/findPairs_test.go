package src

import (
	"fmt"
	"testing"
)

func Test_findPairs(t *testing.T) {
	nums := []int{3, 1, 4, 1, 5}
	fmt.Println(findPairs(nums, 2))
}

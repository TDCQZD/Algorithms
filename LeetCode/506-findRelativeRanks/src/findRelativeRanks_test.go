package src

import (
	"fmt"
	"testing"
)

func Test_findRelativeRanks(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1}

	fmt.Println(findRelativeRanks(nums))
}

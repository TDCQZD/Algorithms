package src

import (
	"fmt"
	"testing"
)

func Test_largestTriangleArea(t *testing.T) {
	nums := [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}

	fmt.Println(largestTriangleArea(nums))
}

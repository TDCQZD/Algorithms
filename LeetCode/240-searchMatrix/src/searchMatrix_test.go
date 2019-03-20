package src

import (
	"fmt"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}
	fmt.Println(searchMatrix(matrix, 4))
}

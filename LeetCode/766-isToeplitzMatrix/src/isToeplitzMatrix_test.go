package src

import (
	"fmt"
	"testing"
)

func Test_isToeplitzMatrix(t *testing.T) {
	data := [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
	fmt.Println(isToeplitzMatrix(data))
}

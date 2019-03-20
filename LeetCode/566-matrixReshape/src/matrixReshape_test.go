package src

import (
	"fmt"
	"testing"
)

func Test_matrixReshape(t *testing.T) {
	nums := [][]int{
		{1, 2},
		{3, 4}}

	fmt.Println(matrixReshape(nums, 1, 4))
	fmt.Println(matrixReshape(nums, 2, 4))

	fmt.Println(matrixReshape2(nums, 1, 4))
	fmt.Println(matrixReshape2(nums, 2, 4))
}

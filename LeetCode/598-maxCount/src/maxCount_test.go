package src

import (
	"fmt"
	"testing"
)

func Test_maxCount(t *testing.T) {
	ope := [][]int{
		{2, 2},
		{3, 3}}

	fmt.Println(maxCount(3, 3, ope))
}

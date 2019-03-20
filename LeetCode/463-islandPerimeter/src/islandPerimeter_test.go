package src

import (
	"fmt"
	"testing"
)

func Test_islandPerimeter(t *testing.T) {
	land := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0}}

	fmt.Println(islandPerimeter(land))

}

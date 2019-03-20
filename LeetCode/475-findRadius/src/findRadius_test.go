package src

import (
	"fmt"
	"testing"
)

func Test_findRadius(t *testing.T) {
	fmt.Println(findRadius([]int{1, 2, 3}, []int{2}))
	fmt.Println(findRadius([]int{1, 2, 3, 4, 99}, []int{1, 4}))
}

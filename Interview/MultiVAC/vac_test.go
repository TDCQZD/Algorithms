package src

import (
	"fmt"
	"testing"
)

func Test_intersection(t *testing.T) {
	num1 := []int{1, 2, 3, 4}
	num2 := []int{2, 3}

	fmt.Println(intersection(num1, num2))
}

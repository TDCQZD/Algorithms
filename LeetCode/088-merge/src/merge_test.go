package src

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	num1 := []int{1, 2, 3, 0, 0, 0}
	num2 := []int{2, 5, 6}
	merge(num1, 3, num2, 3)
	fmt.Println(num1)
}

package src

import (
	"fmt"
	"testing"
)

func TestSingleNumner1(t *testing.T) {
	num1 := []int{2, 2, 1}
	num2 := []int{4, 1, 2, 1, 2}

	res1 := singleNumber(num1)
	res2 := singleNumber(num2)
	fmt.Println(res1, res2)
}

func TestSingleNumner(t *testing.T) {
	var tests = []struct {
		arry []int
		want int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{2}, 2},
		{[]int{2, 2, 2, 2}, 0},
		{nil, 0},
	}

	for _, test := range tests {
		if got := singleNumber(test.arry); got != test.want {
			t.Errorf("singleNumber(%q) = %v", test.arry, got)
		}
	}
}

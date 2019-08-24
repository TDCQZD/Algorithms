package src

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	A := []int{4, 2, 3}
	k := 2
	n := largestSumAfterKNegations(A, k)
	fmt.Println("sum->", n)
}

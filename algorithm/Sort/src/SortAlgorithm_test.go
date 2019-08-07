package src

import (
	"fmt"
	"testing"
)

func TestSortAlgorithm(t *testing.T) {
	// var arrys = []int{24, 69, 80, 57, 30}
	// BulleSort(arrys)

	A := []int{2, 8, 7, 1, 3, 5, 6, 4}
	quickSort(A, 0, 7)
	fmt.Println(A)

	// InsertSort1(A)
}

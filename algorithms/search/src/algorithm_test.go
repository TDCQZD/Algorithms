package src

import (
	"fmt"
	"testing"
)

func TestSortAlgorithm(t *testing.T) {
	testBinarySearch(t)
}

func testBinarySearch(t *testing.T) {
	array := []int{12, 34, 56, 78, 90, 203}
	BinarySearch(array, 0, len(array)-1, 1000)
}

func testQuickSearch(t *testing.T) {
	A := []int{2, 8, 7, 1, 3, 5, 6, 4}
	ret := findKthLargest(A, 0, 7, 8)
	fmt.Println("ret: ", ret)
}

package src

import (
	"fmt"
	"testing"
)

func TestOk(t *testing.T) {
	nums := []int{99, 99}
	flag := containsNearbyDuplicate(nums, 2)
	fmt.Println("result:", flag)
}

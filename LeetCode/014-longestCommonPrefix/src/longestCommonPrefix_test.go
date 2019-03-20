package src

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	// s := []string{"flower", "flow", "flight"}
	s := []string{"dog", "racecar", "car"}
	res := longestCommonPrefix(s)
	fmt.Println(res)
}

package src

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	// length := lengthOfLongestSubstring("abcabcbb")
	// fmt.Println(length)

	str, length := lengthOfLongestSubstringInt("abcabcbb")
	fmt.Println(str,length)
}

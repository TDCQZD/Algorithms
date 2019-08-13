package src

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	s := "Hello"
	s = reverseString(s)
	fmt.Println(s)
}

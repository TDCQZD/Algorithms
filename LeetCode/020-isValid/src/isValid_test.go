package src

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	s := "([)]"
	res := isValid(s)
	fmt.Println(res)
}

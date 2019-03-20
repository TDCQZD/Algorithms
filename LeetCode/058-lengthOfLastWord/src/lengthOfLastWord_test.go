package src

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	str := "Hello World"
	res := lengthOfLastWord(str)
	fmt.Println(res)
	res = lengthOfLastWord2(str)
	fmt.Println(res)
}

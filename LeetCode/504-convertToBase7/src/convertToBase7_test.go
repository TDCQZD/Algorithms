package src

import (
	"fmt"
	"testing"
)

func Test_convertToBase7(t *testing.T) {
	fmt.Println(convertToBase7(7))
	fmt.Println(convertToBase7(-7))
	fmt.Println(convertToBase7(100))
}

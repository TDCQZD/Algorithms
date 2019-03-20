package src

import (
	"fmt"
	"testing"
)

func Test_uniqueMorseRepresentations(t *testing.T) {
	strList := []string{"abc", "bcd"}
	fmt.Println(uniqueMorseRepresentations(strList))
}

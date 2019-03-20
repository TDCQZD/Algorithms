package src

import (
	"fmt"
	"testing"
)

func Test_toGoatLatin(t *testing.T) {
	S := "I speak Goat Latin"

	fmt.Println(toGoatLatin(S))
}

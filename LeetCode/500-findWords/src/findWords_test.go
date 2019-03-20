package src

import (
	"fmt"
	"testing"
)

func Test_findWords(t *testing.T) {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Println(findWords(words))
}

package src

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	stack := New(1, 2, 3, 4)
	fmt.Println(stack.Values())
	stack.Push(5, 6)
	fmt.Println(stack.Values())
	fmt.Println(stack.String())
	stack.Pop()
	fmt.Println(stack.Values())
	value, flag := stack.Get(3)
	if flag {
		fmt.Println(value)
	}
	fmt.Println(stack.Values())
	stack.Remove(1)
	fmt.Println(stack.Values())
}

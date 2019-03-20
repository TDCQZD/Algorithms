package src

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	queue := New(1, 2, 3, 4)
	fmt.Println(queue.Values())
	queue.Push(5, 6)
	fmt.Println(queue.Values())
	fmt.Println(queue.String())
	queue.Pop()
	fmt.Println(queue.Values())
	value, flag := queue.Get(3)
	if flag {
		fmt.Println(value)
	}
	fmt.Println(queue.Values())
	queue.Remove(1)
	fmt.Println(queue.Values())
}

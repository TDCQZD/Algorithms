package src

import (
	"fmt"
	"testing"
)

func TestLink(t *testing.T) {
	list := New(1, 2)
	fmt.Println(list)
	fmt.Println(list.Values())
	list.Add(3, 4)
	fmt.Println(list.Values())
	list.Append(5)
	fmt.Println(list.Values())
	fmt.Println("IndexOf:", list.IndexOf(3))
	fmt.Println(list.Get(0))
	list.Prepend(6)
	fmt.Println(list.Values())
	fmt.Println("Contains:", list.Contains(3))

	list.Prepend(7, 8)
	fmt.Println(list.Values())
	list.Swap(0, 1)
	fmt.Println(list.Values())
	list.Insert(0,10)
	fmt.Println(list.Values())
	list.Set(1,11)
	fmt.Println(list.Values())
	fmt.Println(list.String())

	res:= list.Remove(0)
	fmt.Println(list.Values())
	fmt.Println("remove element:",res)
	fmt.Println("Empty:", list.Empty())
	fmt.Println("Size:", list.ListSize())
	list.Clear()
	fmt.Println("Empty:", list.Empty())
	fmt.Println(list.Values())
}

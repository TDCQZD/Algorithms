package src

import (
	"fmt"
	"strings"
)

type Node struct {
	Val  interface{}
	Next *Node
}
type List struct {
	Head *Node
	End  *Node
	Size int
}

func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *List) Add(values ...interface{}) {
	for _, value := range values {
		newNode := &Node{Val: value}
		if list.Size == 0 {
			list.Head = newNode
			list.End = newNode
		} else {
			list.End.Next = newNode
			list.End = newNode
		}
		list.Size++
	}
}
func (list *List) Append(values ...interface{}) {
	list.Add(values...)
}

// 添加到链表头
func (list *List) Prepend(values ...interface{}) {
	// in reverse to keep passed order i.e. ["c","d"] -> Prepend(["a","b"]) -> ["a","b","c",d"]
	for v := len(values) - 1; v >= 0; v-- {
		newElement := &Node{Val: values[v], Next: list.Head}
		list.Head = newElement
		if list.Size == 0 {
			list.End = newElement
		}
		list.Size++
	}
}

func (list *List) Get(index int) (interface{}, bool) {

	if !list.withinRange(index) {
		return nil, false
	}

	element := list.Head
	for e := 0; e != index; e, element = e+1, element.Next {
	}

	return element.Val, true
}

func (list *List) Remove(index int) (value interface{}) {
	value = nil
	if !list.withinRange(index) {
		return
	}

	if list.Size == 1 {
		list.Clear()
		return
	}

	var beforeElement *Node
	element := list.Head
	for e := 0; e != index; e, element = e+1, element.Next {
		beforeElement = element
	}

	if element == list.Head {
		list.Head = element.Next
	}
	if element == list.End {
		list.End = beforeElement
	}
	if beforeElement != nil {
		beforeElement.Next = element.Next
	}
	value = element
	element = nil

	list.Size--
	return
}

func (list *List) Contains(values ...interface{}) bool {

	if len(values) == 0 {
		return true
	}
	if list.Size == 0 {
		return false
	}
	for _, value := range values {
		found := false
		for element := list.Head; element != nil; element = element.Next {
			if element.Val == value {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (list *List) Values() []interface{} {
	values := make([]interface{}, list.Size, list.Size)
	for e, element := 0, list.Head; element != nil; e, element = e+1, element.Next {
		values[e] = element.Val
	}
	return values
}
func (list *List) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) && i != j {
		var element1, element2 *Node
		for e, currentElement := 0, list.Head; element1 == nil || element2 == nil; e, currentElement = e+1, currentElement.Next {
			switch e {
			case i:
				element1 = currentElement
			case j:
				element2 = currentElement
			}
		}
		element1.Val, element2.Val = element2.Val, element1.Val
	}
}

func (list *List) Insert(index int, values ...interface{}) {

	if !list.withinRange(index) {
		// Append
		if index == list.Size {
			list.Add(values...)
		}
		return
	}

	list.Size += len(values)

	var beforeElement *Node
	foundElement := list.Head
	for e := 0; e != index; e, foundElement = e+1, foundElement.Next {
		beforeElement = foundElement
	}

	if foundElement == list.Head {
		oldNextElement := list.Head
		for i, value := range values {
			newElement := &Node{Val: value}
			if i == 0 {
				list.Head = newElement
			} else {
				beforeElement.Next = newElement
			}
			beforeElement = newElement
		}
		beforeElement.Next = oldNextElement
	} else {
		oldNextElement := beforeElement.Next
		for _, value := range values {
			newElement := &Node{Val: value}
			beforeElement.Next = newElement
			beforeElement = newElement
		}
		beforeElement.Next = oldNextElement
	}
}

func (list *List) Set(index int, value interface{}) {

	if !list.withinRange(index) {
		// Append
		if index == list.Size {
			list.Add(value)
		}
		return
	}

	foundElement := list.Head
	for e := 0; e != index; {
		e, foundElement = e+1, foundElement.Next
	}
	foundElement.Val = value
}

func (list *List) IndexOf(value interface{}) int {
	if list.Size == 0 {
		return -1
	}
	for index, element := range list.Values() {
		if element == value {
			return index
		}
	}
	return -1
}

func (list *List) Empty() bool {
	return list.Size == 0
}

func (list *List) ListSize() int {
	return list.Size
}

func (list *List) Clear() {
	list.Size = 0
	list.Head = nil
	list.End = nil
}

func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.Size
}

func (list *List) String() string {
	str := "SinglyLinkedList\n"
	values := []string{}
	for element := list.Head; element != nil; element = element.Next {
		values = append(values, fmt.Sprintf("%v", element.Val))
	}
	str += strings.Join(values, ", ")
	return str
}

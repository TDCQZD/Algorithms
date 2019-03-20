package src

import (
	"fmt"
	"strings"
)

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

type ArrayStack struct {
	elements []interface{}
	size     int
}

// New instantiates a new empty stack
func New(values ...interface{}) *ArrayStack {
	stack := &ArrayStack{}
	if len(values) > 0 {
		stack.Push(values...)
	}
	return stack
}

// Push adds a value onto the top of the stack
func (stack *ArrayStack) Push(values ...interface{}) {
	stack.growBy(len(values))
	for _, value := range values {
		stack.elements[stack.size] = value
		stack.size++
	}
}

// Pop removes top element on stack and returns it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to pop.
func (stack *ArrayStack) Pop() (value interface{}, ok bool) {

	value, ok = stack.Get(stack.Size() - 1)
	stack.Remove(stack.Size() - 1)
	return
}
func (stack *ArrayStack) Remove(index int) {

	if !stack.withinRange(index) {
		return
	}

	stack.elements[index] = nil                                      // cleanup reference
	copy(stack.elements[index:], stack.elements[index+1:stack.size]) // shift to the left by one (slow operation, need ways to optimize this)
	stack.size--

	stack.shrink()
}

func (stack *ArrayStack) Get(index int) (interface{}, bool) {

	if !stack.withinRange(index) {
		return nil, false
	}
	return stack.elements[index], true
}

// Peek returns top element on the stack without removing it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to peek.
func (stack *ArrayStack) Peek() (value interface{}, ok bool) {
	return stack.Get(stack.Size() - 1)
}

// Empty returns true if stack does not contain any elements.
func (stack *ArrayStack) Empty() bool {
	return stack.size == 0
}

// Size returns number of elements within the stack.
func (stack *ArrayStack) Size() int {
	return stack.size
}

// Clear removes all elements from the stack.
func (stack *ArrayStack) Clear() {
	stack.size = 0
	stack.elements = []interface{}{}
}

// Values returns all elements in the stack (LIFO order).
func (stack *ArrayStack) Values() []interface{} {
	size := stack.Size()
	elements := make([]interface{}, size, size)
	for i := 1; i <= size; i++ {
		elements[size-i], _ = stack.Get(i - 1) // in reverse (LIFO)
	}
	return elements
}

// String returns a string representation of container
func (stack *ArrayStack) String() string {
	str := "ArrayStack to string :"
	values := []string{}
	for _, value := range stack.elements {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is within bounds of the element
func (stack *ArrayStack) withinRange(index int) bool {
	return index >= 0 && index < stack.Size()
}

func (stack *ArrayStack) growBy(n int) {
	// When capacity is reached, grow by a factor of growthFactor and add number of elements
	currentCapacity := cap(stack.elements)
	if stack.size+n >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		stack.resize(newCapacity)
	}
}

func (stack *ArrayStack) resize(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, stack.elements)
	stack.elements = newElements
}

// Shrink the array if necessary, i.e. when size is shrinkFactor percent of current capacity
func (stack *ArrayStack) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	// Shrink when size is at shrinkFactor * capacity
	currentCapacity := cap(stack.elements)
	if stack.size <= int(float32(currentCapacity)*shrinkFactor) {
		stack.resize(stack.size)
	}
}

package src

import (
	"fmt"
	"strings"
)

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

type ArrayQueue struct {
	elements []interface{}
	size     int
}

func New(values ...interface{}) *ArrayQueue {
	queue := &ArrayQueue{}
	if len(values) > 0 {
		queue.Push(values...)
	}
	return queue
}

func (queue *ArrayQueue) Push(values ...interface{}) {
	queue.growBy(len(values))
	for _, value := range values {
		queue.elements[queue.size] = value
		queue.size++
	}
}

func (queue *ArrayQueue) Pop() (value interface{}, ok bool) {

	// value, ok = queue.Get(queue.Size() - 1)
	// queue.Remove(queue.Size() - 1)
	value, ok = queue.Get(0)
	queue.Remove(0)
	return
}
func (queue *ArrayQueue) Remove(index int) {

	if !queue.withinRange(index) {
		return
	}

	queue.elements[index] = nil                                      // cleanup reference
	copy(queue.elements[index:], queue.elements[index+1:queue.size]) // shift to the left by one (slow operation, need ways to optimize this)
	queue.size--

	queue.shrink()
}

func (queue *ArrayQueue) Get(index int) (interface{}, bool) {

	if !queue.withinRange(index) {
		return nil, false
	}
	return queue.elements[index], true
}

func (queue *ArrayQueue) Peek() (value interface{}, ok bool) {
	// return queue.Get(queue.Size() - 1)
	return queue.Get(0)
}

func (queue *ArrayQueue) Empty() bool {
	return queue.size == 0
}

func (queue *ArrayQueue) Size() int {
	return queue.size
}

func (queue *ArrayQueue) Clear() {
	queue.size = 0
	queue.elements = []interface{}{}
}

func (queue *ArrayQueue) Values() []interface{} {
	size := queue.Size()
	elements := make([]interface{}, size, size)
	for i := 1; i <= size; i++ {
		elements[i-1], _ = queue.Get(i - 1)
	}
	return elements
}

func (queue *ArrayQueue) String() string {
	str := "ArrayQueue to string :"
	values := []string{}
	for _, value := range queue.elements {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

func (queue *ArrayQueue) withinRange(index int) bool {
	return index >= 0 && index < queue.Size()
}

func (queue *ArrayQueue) growBy(n int) {
	// When capacity is reached, grow by a factor of growthFactor and add number of elements
	currentCapacity := cap(queue.elements)
	if queue.size+n >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		queue.resize(newCapacity)
	}
}

func (queue *ArrayQueue) resize(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, queue.elements)
	queue.elements = newElements
}

// Shrink the array if necessary, i.e. when size is shrinkFactor percent of current capacity
func (queue *ArrayQueue) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	// Shrink when size is at shrinkFactor * capacity
	currentCapacity := cap(queue.elements)
	if queue.size <= int(float32(currentCapacity)*shrinkFactor) {
		queue.resize(queue.size)
	}
}

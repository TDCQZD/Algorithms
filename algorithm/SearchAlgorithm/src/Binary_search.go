package src

import (
	"fmt"
)

func BinarySearch(array []int, startIndex, endIndex, find int) {
	if len(array) == 0 {
		return
	}
	if startIndex > endIndex {
		fmt.Printf("%d 不是数组中元素!!!!\n", find)
		return
	}
	middle := (startIndex + endIndex) / 2

	if array[middle] > find {
		BinarySearch(array, startIndex, middle-1, find)
	} else if array[middle] < find {
		BinarySearch(array, middle+1, endIndex, find)
	} else {
		fmt.Printf("%d 在数组中的下标为 %d \n", find, middle)
	}
}

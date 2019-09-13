package src

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		reverse(A[i])
		invert(A[i])
	}
	return A
}

func reverse(array []int) {
	for left, right := 0, len(array)-1; left < right; {
		array[left], array[right] = array[right], array[left]
		left++
		right--
	}
}

func invert(array []int) {
	for i := 0; i < len(array); i++ {
		if 1 == array[i] {
			array[i] = 0
		} else if 0 == array[i] {
			array[i] = 1
		}
	}
}
func flipAndInvertImage1(A [][]int) [][]int {
	C := len(A[0])
	for _, row := range A {
		for i := 0; i < (C+1)/2; i++ {
			row[i], row[C-1-i] = invert10(row[C-1-i]), invert10(row[i])
		}
	}
	return A
}
func invert10(array int) int {
	if 1 == array {
		array = 0
	} else if 0 == array {
		array = 1
	}
	return array
}

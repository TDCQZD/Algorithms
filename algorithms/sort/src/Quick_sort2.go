package src

import "fmt"

func partition(A []int, p, r int) int {
	x := A[r]
	i := p - 1
	j := p
	for j >= p && j < r {
		if A[j] <= x {
			i++
			A[i], A[j] = A[j], A[i]
		}
		fmt.Println(i, j, A)
		j++
	}
	A[i+1], A[r] = A[r], A[i+1]
	fmt.Println(A, p, r)
	return i + 1
}

func quickSort(A []int, p, r int) {
	if p < r {
		q := partition(A, p, r)
		quickSort(A, p, q-1)
		quickSort(A, q+1, r)
	}
}

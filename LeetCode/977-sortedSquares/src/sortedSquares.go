package src

import (
	"sort"
)

func sortedSquares(A []int) []int {
	N := len(A)
	j := 0
	for j < N && A[j] < 0 {
		j++
	}

	i := j - 1

	ans := make([]int, len(A))
	t := 0

	for i >= 0 && j < N {
		if A[i]*A[i] < A[j]*A[j] {
			ans[t] = A[i] * A[i]
			t++
			i--
		} else {
			ans[t] = A[j] * A[j]
			t++
			j++
		}
	}

	for i >= 0 {
		ans[t] = A[i] * A[i]
		t++
		i--
	}
	for j < N {
		ans[t] = A[j] * A[j]
		t++
		j++
	}

	return ans
}
func sortedSquares1(A []int) []int {
	ans := make([]int, len(A))

	for i := 0; i < len(A); i++ {
		ans[i] = A[i] * A[i]
	}
	sort.Ints(ans)
	return ans
}

package src

func isMonotonic(A []int) bool {
	if len(A) == 1 {
		return true
	}
	a, b := 0, 0
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			a = 1
		}
		if A[i] < A[i+1] {
			b = 1
		}
	}
	if a+b == 2 {
		return false
	} else {
		return true
	}

}

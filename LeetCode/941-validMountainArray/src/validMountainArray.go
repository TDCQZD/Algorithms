package src

func validMountainArray(A []int) bool {
	N := len(A)
	i := 0

	// walk up
	for i+1 < N && A[i] < A[i+1] {
		i++
	}

	// peak can't be first or last
	if i == 0 || i == N-1 {
		return false
	}

	// walk down
	for i+1 < N && A[i] > A[i+1] {
		i++
	}

	return i == N-1

}

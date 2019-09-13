package src

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	S := 0
	for _, x := range A {
		if x%2 == 0 {
			S += x
		}
	}

	ans := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		val := queries[i][0]
		index := queries[i][1]
		if A[index]%2 == 0 {
			S -= A[index]
		}
		A[index] += val
		if A[index]%2 == 0 {
			S += A[index]
		}
		ans[i] = S
	}

	return ans
}

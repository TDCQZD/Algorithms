package src

func sortArrayByParityII(A []int) []int {
	j := 1
	for i := 0; i < len(A); i += 2 {
		if A[i]%2 == 1 {
			for A[j]%2 == 1 {
				j += 2
			}
			// Swap A[i] and A[j]
			A[i], A[j] = A[j], A[i]
		}
	}
	return A

}
func sortArrayByParityII1(A []int) []int {
	N := len(A)
	ans := make([]int, N)

	t := 0
	for _, x := range A {
		if x%2 == 0 {
			ans[t] = x
			t += 2
		}
	}

	t = 1
	for _, x := range A {
		if x%2 == 1 {
			ans[t] = x
			t += 2
		}
	}
	return ans

}

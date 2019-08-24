package src

func minDeletionSize(A []string) int {
	ans := 0
	for i := 0; i < len(A[0]); i++ {
		for j := 0; j < len(A)-1; j++ {
			if A[j][i:i+1] > A[j+1][i:i+1] {
				ans++
				break
			}
		}
	}
	return ans
}

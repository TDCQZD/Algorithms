package src

func superEggDrop(K int, N int) int {
    moves := 0
	dp := make([]int, K+1) // 1 <= K <= 100
	for dp[K] < N {
		for i := K; i > 0; i-- {		
			dp[i] += dp[i-1] + 1			
		}
		moves++
	}
	return moves
}
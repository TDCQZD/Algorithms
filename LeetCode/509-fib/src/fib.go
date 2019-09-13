package src

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N <= 2 {
		return 1
	}
	n1, n2 := 1, 1 // n1为n-1，n2为n-2
	for i := 3; i < N; i++ {
		n1, n2 = n1+n2, n1
	}
	return n1 + n2
}
func fib1(N int) int {
	if N == 1 || N == 2 {
		return 1
	}
	if N > 2 {
		return fib1(N-1) + fib1(N-2)
	}
	return 0
}
func Fib2(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	array := make([]int, n+1)
	array[0] = 0
	array[1] = 1
	i := 2
	for {
		array[i] = array[i-1] + array[i-2]
		i++
		if i > n {
			break
		}
	}

	return array[n]
}

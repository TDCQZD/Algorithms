package test

/*
题目：翻书问题或者跳台阶问题。
每次可以翻1页书或者2页书，一本N页的书共有多少中翻法。
*/

/*FibonacciByRecursive 递归实现——不推荐解法 */
func FibonacciByRecursive(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if n > 1 {
		return FibonacciByRecursive(n-1) + FibonacciByRecursive(n-2)
	}
	return 0
}

/*FibonacciBycycle 循环实现——推荐解法*/
func FibonacciBycycle(n int) int {
	array := make([]int, n+1)

	array[0] = 1
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

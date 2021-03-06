package src

func partition(A []int, p, r int) int {
	x := A[r]
	i := p - 1
	j := p
	for j >= p && j < r {
		if A[j] >= x {
			i++
			A[i], A[j] = A[j], A[i]
		}
		j++
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}

// 快速选择
func findKthLargest(A []int, p, r, k int) int {
	if p <= r {
		q := partition(A, p, r)
		if k-1 == q {
			return A[q]
		} else if k-1 < q {
			return findKthLargest(A, p, q-1, k)
		} else {
			return findKthLargest(A, q+1, r, k)
		}
	}
	return -1000000
}

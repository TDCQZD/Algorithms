package src

func fairCandySwap(A []int, B []int) []int {
	sa, sb := 0, 0            // sum of A, B respectively
	setB := make(map[int]int) // If Alice gives x, she expects to receive x + delta
	for _, x := range A {
		sa += x
	}
	for _, y := range B {
		sb += y
		setB[y] = 0
	}

	delta := (sb - sa) / 2

	for _, x := range A {
		if _, ok := setB[x+delta]; ok {
			return []int{x, x + delta}
		}
	}
	return nil
}

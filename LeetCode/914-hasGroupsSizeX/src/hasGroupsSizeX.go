package src

func hasGroupsSizeX(deck []int) bool {
	if len(deck) < 2 {
		return false
	}
	xmap := make(map[int]int)

	for i := 0; i < len(deck); i++ {
		xmap[deck[i]]++
	}
	g := -1
	for _, v := range xmap {
		if v > 0 {
			if g == -1 {
				g = v
			} else {
				g = gcd(g, v)
			}
		}
	}
	return g >= 2

}

func gcd(x, y int) int {
	if x == 0 {
		return y
	} else {
		return gcd(y%x, x)
	}
}

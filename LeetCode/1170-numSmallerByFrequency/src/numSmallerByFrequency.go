package src

func numSmallerByFrequency(queries []string, words []string) []int {
	q := map[int]int{}
	for i := 0; i < len(queries); i++ {
		q[i] = f(queries[i])
	}
	w := [11]int{}
	for i := 0; i < len(words); i++ {
		w[f(words[i])]++
	}
	ret := make([]int, len(q))
	for k, v := range q {
		for i := v + 1; i < len(w); i++ {
			ret[k] += w[i]
		}
	}
	return ret
}
func f(str string) int {
	bs := [26]int{}
	for i := 0; i < len(str); i++ {
		bs[int(str[i]-'a')]++
	}
	for i := 0; i < len(bs); i++ {
		if bs[i] > 0 {
			return bs[i]
		}
	}
	return 0
}

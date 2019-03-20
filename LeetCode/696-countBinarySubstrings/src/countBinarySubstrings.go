package src

func countBinarySubstrings(s string) int {
	prevRun, curRun, ans := 0, 1, 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			curRun++
		} else {
			prevRun = curRun
			curRun = 1
		}
		if prevRun >= curRun {
			ans++
		}
	}
	return ans
}

package src

func maxArea(height []int) int {
	maxarea := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			maxarea = max(maxarea, min(height[i], height[j])*(j-i))
		}
	}

	return maxarea
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
func maxArea1(height []int) int {

	maxarea, l, r := 0, 0, len(height)-1
	for l < r {
		maxarea = max(maxarea, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxarea
}

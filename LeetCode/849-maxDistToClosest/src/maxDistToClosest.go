package src

func maxDistToClosest(seats []int) int {
	count1 := 0
	count2 := 0
	i, j := 0, len(seats)-1
	// count1记录开头连续0的个数
	for seats[i] == 0 {
		count1++
		i++
	}
	// count2记录结尾连续0的个数
	for seats[j] == 0 {
		count2++
		j--
	}
	// countmax记录从第一个1到最后一个1之间，连续0的最大值
	countmid, countmax := 0, 0
	for k := i + 1; k <= j; k++ {
		if seats[k] == 0 {
			countmid++
		} else {
			countmax = max(countmax, countmid)
			countmid = 0
		}
	}
	// 返回count1, count2, (countmax+1)/2三者中最大值
	return max(max(count1, count2), (countmax+1)/2)

}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

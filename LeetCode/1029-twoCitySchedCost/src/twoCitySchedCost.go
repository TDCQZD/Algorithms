package src

import "sort"

// 假设所有人员到A城，排序，减去一半人的额外消耗
func twoCitySchedCost(costs [][]int) int {
	t := len(costs)
	cha := []int{}
	sum := 0
	for i := 0; i < t; i++ {
		cha = append(cha, costs[i][1]-costs[i][0]) //计算人员到A城与到B城的消耗差
		sum += costs[i][0]                         //加上所有人到A城的费用
	}
	sort.Ints(cha)

	for j := 0; j < t/2; j++ {
		sum += cha[j] //减去应到B城的人员的额外消耗
	}

	return sum

}
func twoCitySchedCost2(costs [][]int) int {

	// 二维数组排序

	total := 0
	n := len(costs) / 2
	for i := 0; i < n; i++ {
		total += costs[i][0] + costs[i+n][1]
	}
	return total
}

// 二维数组排序
func TwoArraySort(arrys [][]int) [][]int {

	return nil
}

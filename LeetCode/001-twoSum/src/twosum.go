package src

//  方法一：暴力法
func TwoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 方法二：两遍哈希表
func TwoSum2(nums []int, target int) []int {
	mapSum := make(map[int]int, len(nums))
	for k, v := range nums {
		mapSum[v] = k
	}
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if _, ok := mapSum[diff]; ok && mapSum[diff] != i {
			return []int{i,mapSum[diff]}
		}
	return nil
}


// 方法三：一遍哈希表
func TwoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))
	// 循环遍历元素,使用map来存储元素和下标之间的映射关系
	for indexB, b := range nums {
		//  判断map的key是否存在
		if indexA, ok := index[target-b]; ok {
			return []int{indexA, indexB}
		}
		// 使用map来存储元素和下标之间的映射关系
		index[b] = indexB
	}
	return nil
}
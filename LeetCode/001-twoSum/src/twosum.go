package src

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


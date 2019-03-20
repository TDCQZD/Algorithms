package src

func TwoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))

	for indexB, b := range nums {
		if indexA, ok := index[target-b]; ok {
			return []int{indexA, indexB}
		}
		index[b] = indexB
	}
	return nil
}

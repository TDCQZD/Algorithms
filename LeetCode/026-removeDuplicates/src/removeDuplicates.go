package src

func removeDuplicates1(nums []int) (newLength int, newNums []int) {
	if nil == nums || len(nums) < 1 {
		return 0, nil
	}

	newLength = 1
	newNums = append(newNums, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[newLength-1] {
			nums[newLength] = nums[i]
			newNums = append(newNums, nums[i])
			newLength++
		}
	}
	return newLength, newNums
}

func removeDuplicates(nums []int) int {
	if nil == nums || len(nums) < 1 {
		return 0
	}

	newLength := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[newLength-1] {
			nums[newLength] = nums[i]
			newLength++
		}

	}
	return newLength
}

// 双指针
func removeDuplicates2(nums []int) int {
	if nil == nums || len(nums) < 1 {
		return 0
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}

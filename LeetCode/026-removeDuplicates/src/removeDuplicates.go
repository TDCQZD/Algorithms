package src

func removeDuplicates(nums []int) (newLength int, newNums []int) {
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

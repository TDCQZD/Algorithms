package src

func removeElement(nums []int, val int) (idx int, newNums []int) {
	if nil == nums || len(nums) < 1 {
		return 0, nil
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[idx] = nums[i]
			newNums = append(newNums, nums[i])
			idx++
		}
	}
	return idx, newNums
}
func removeElement1(nums []int, val int) int {
	if nil == nums || len(nums) < 1 {
		return 0
	}

	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}

package src

// 一次遍历 三路快速排序方法
func sortColors(nums []int) {
	// 对于所有 idx < i : nums[idx < i] = 0
	// j是当前考虑元素的下标
	curr := 0
	p0 := 0
	// 对于所有 idx > k : nums[idx > k] = 2
	p2 := len(nums) - 1

	tmp := 0
	for curr <= p2 {
		if nums[curr] == 0 {
			// 交换第 p0个和第curr个元素
			// i++，j++
			tmp = nums[p0]
			nums[p0] = nums[curr]
			nums[curr] = tmp

			p0++
			curr++
		} else if nums[curr] == 2 {
			// 交换第k个和第curr个元素
			// p2--
			tmp = nums[curr]
			nums[curr] = nums[p2]
			nums[p2] = tmp
			p2--
		} else {
			curr++
		}

	}
}


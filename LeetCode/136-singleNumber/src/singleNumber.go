package src

func singleNumber(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	ans := 0
	for _, value := range nums {
		ans ^= value //0 ^ n = n 异或：相同为0，不同为1. 异或同一个数两次，原数不变。
	}
	return ans

}

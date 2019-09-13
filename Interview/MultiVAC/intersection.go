package src

func intersection(nums1 []int, nums2 []int) []int {
	dict := make(map[int]int)
	for _, v := range nums1 {
		dict[v] = 1
	}

	for _, v := range nums2 {
		if 1 == dict[v] {
			dict[v] = 2
		}
	}

	ans := make([]int, 0)
	for k, v := range dict {
		if 2 == v {
			ans = append(ans, k)
		}
	}
	return ans
}

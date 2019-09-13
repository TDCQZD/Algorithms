package src

import "sort"

// 双指针 尾插法
func merge(nums1 []int, m int, nums2 []int, n int) {
	idm, idn := m-1, n-1
	for i := m + n - 1; i >= 0; i-- {
		if idm < 0 {
			nums1[i] = nums2[idn]
			idn--
			continue
		}
		if idn < 0 {
			nums1[i] = nums1[idm]
			idm--
			continue
		}
		if nums1[idm] >= nums2[idn] {
			nums1[i] = nums1[idm]
			idm--
			continue
		}
		if nums1[idm] < nums2[idn] {
			nums1[i] = nums2[idn]
			idn--
			continue
		}
	}
}

// 合并后排序
func merge2(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[:m]
	for i := 0; i < n; i++ {
		nums1 = append(nums1, nums2[i])
	}
	sort.Ints(nums1)
}

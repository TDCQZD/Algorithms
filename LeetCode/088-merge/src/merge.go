package src

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
func merge2(nums1 []int, m int, nums2 []int, n int) {
	nums1_copy := nums1
	p1, p2, p := 0, 0, 0
	for p1 < m && p2 < n {
		p++
		if nums1_copy[p1] < nums2[p2] {
			nums1[p] = nums1_copy[p1+1]
			p1++
		} else {
			nums1[p] = nums2[p2+1]
			p2++
		}

	}
	if p1 < m {
		for i := p1 + p2; i <= m+n-p1-p2; i++ {
			nums1[i] = nums1_copy[p1]
			p1++
		}
	}
	if p2 < n {
		for i := p1 + p2; i <= m+n-p1-p2; i++ {
			nums1[i] = nums2[p2]
			p2++
		}
	}

}
func merge3(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p := m-1, n-1, m+n-1
	for p1 >= 0 && p2 > 0 {

		if nums1[p1] < nums2[p2] {
			nums1[p] = nums2[p2]
			p2--
		} else {
			nums1[p] = nums1[p1]
			p1--
		}
		p--
	}
	
}

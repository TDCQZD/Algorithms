package src

func moveZeroes(nums []int) {
	if nil == nums || len(nums) == 0 {
		return
	}

	insertPos := 0
	for _, values := range nums {
		if 0 != values {
			nums[insertPos] = values
			insertPos++
		}
	}

	for insertPos < len(nums) {
		nums[insertPos] = 0
		insertPos++
	}
}
func moveZeroes1(nums []int) {
	if nil == nums || len(nums) <= 1 {
		return
	}
	p1 := 0
	p2 := 1
	for {
		if nums[p1] == 0 && nums[p2] == 0 {
			p2++
			if p2 > len(nums)-1 {
				break
			} else {
				continue
			}
		}
		if nums[p1] == 0 && nums[p2] != 0 {
			nums[p1], nums[p2] = nums[p2], nums[p1]
			p1++
			p2++
			if p2 > len(nums)-1 {
				break
			} else {
				continue
			}
		}
		if nums[p1] != 0 && nums[p2] == 0 {
			p1++
			p2++
			if p2 > len(nums)-1 {
				break
			} else {
				continue
			}
		}
		if nums[p1] != 0 && nums[p2] != 0 {
			p1 = p2 + 1
			p2 = p1 + 1
			if p2 > len(nums)-1 {
				break
			} else {
				continue
			}
		}

	}
}

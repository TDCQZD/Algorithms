package src

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			five++
		} else if bills[i] == 10 {
			if five == 0 {
				return false
			}
			ten++
			five--
		} else if bills[i] == 20 {
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}

		}
	}

	return true
}

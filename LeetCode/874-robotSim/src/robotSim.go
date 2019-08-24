package src

func robotSim(commands []int, obstacles [][]int) int {
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	x, y, di, ans := 0, 0, 0, 0

	// obstacleSlice := make([]int64, 0)
	obstacleMap := make(map[int64]int64)
	for _, obstacle := range obstacles {
		ox := int64(obstacle[0] + 30000)
		oy := int64(obstacle[1] + 30000)
		// obstacleSlice = append(obstacleSlice, (ox<<16)+oy)
		obstacleMap[(ox<<16)+oy] = (ox << 16) + oy
	}

	for _, cmd := range commands {
		if cmd == -1 {
			di = (di + 1) % 4
		} else if cmd == -2 {
			di = (di + 3) % 4
		} else {
			for i := 0; i < cmd; i++ {
				nx := x + dx[di]
				ny := y + dy[di]
				code := ((int64(nx) + 30000) << 16) + (int64(ny) + 30000)
				// if !IsExistArray(obstacleSlice, code) {
				// 	x = nx
				// 	y = ny
				// 	ans = Max(ans, x*x+y*y)
				// }
				_, ok := obstacleMap[code]
				if !ok {
					x = nx
					y = ny
					ans = Max(ans, x*x+y*y)
				}

			}
		}
	}
	return ans
}

func IsExistArray(arr []int64, value int64) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

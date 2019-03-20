package src

func searchMatrix(matrix [][]int, target int) bool {

	if len(matrix) > 0 && len(matrix[0]) > 0 {

		i := 0
		j := len(matrix[0]) - 1
		for i < len(matrix) && j >= 0 {
			if matrix[i][j] == target {
				return true
			} else if matrix[i][j] > target {
				j -= 1
			} else {
				i += 1
			}
		}
		return false
	} else {
		return false
	}

}

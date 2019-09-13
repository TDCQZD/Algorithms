package src

// 无需考虑复杂，直接判断是否符合要求即可
// 只需判断每个元素是否和左上角的元素相等就ok了
func isToeplitzMatrix(matrix [][]int) bool {
	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[0])-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true
}

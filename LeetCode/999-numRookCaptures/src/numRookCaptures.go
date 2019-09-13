package src

import "fmt"

func numRookCaptures(board [][]byte) int {
	var x, y int
	var A = board
	for i, _ := range A {
		for j, v := range A[i] {
			if v == 'R' {
				x, y = i, j
				goto LABLE
			}
		}
	}
LABLE:
	fmt.Println(x, y)
	var res = 0
	for i := x + 1; i < 8; i++ {
		if A[i][y] == 'B' {
			break
		}
		if A[i][y] == 'p' {
			res++
			break
		}

	}
	for i := x - 1; i >= 0; i-- {
		if A[i][y] == 'B' {
			break
		}
		if A[i][y] == 'p' {
			res++
			break
		}

	}
	for i := y + 1; i < 8; i++ {
		if A[x][i] == 'B' {
			break
		}
		if A[x][i] == 'p' {
			res++
			break
		}

	}
	for i := y - 1; i >= 0; i-- {
		if A[x][i] == 'B' {
			break
		}
		if A[x][i] == 'p' {
			res++
			break
		}

	}
	return res
}
func numRookCaptures1(board [][]byte) int {
	rx, ry := -1, -1
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'R' {
				rx, ry = i, j
				break
			}
		}
		if rx != -1 {
			break
		}
	}
	count := 0
	for i := rx - 1; i >= 0; i-- {
		if board[i][ry] == 'B' {
			break
		} else if board[i][ry] == 'p' {
			count++
			break
		}
	}
	for i := rx + 1; i < len(board); i++ {
		if board[i][ry] == 'B' {
			break
		} else if board[i][ry] == 'p' {
			count++
			break
		}
	}
	for j := ry - 1; j >= 0; j-- {
		if board[rx][j] == 'B' {
			break
		} else if board[rx][j] == 'p' {
			count++
			break
		}
	}
	for j := ry + 1; j < len(board[rx]); j++ {
		if board[rx][j] == 'B' {
			break
		} else if board[rx][j] == 'p' {
			count++
			break
		}
	}
	return count
}

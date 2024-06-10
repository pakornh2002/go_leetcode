package isvalidsudoku

func IsValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		vertical := make(map[byte]int)
		horizontal := make(map[byte]int)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				vertical[board[i][j]]++
				if vertical[board[i][j]] > 1 {
					return false
				}
			}
			if board[j][i] != '.' {
				horizontal[board[j][i]]++
				if horizontal[board[j][i]] > 1 {
					return false
				}
			}
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			grid := make(map[byte]int)
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if board[i+k][j+l] != '.' {
						grid[board[i+k][j+l]]++
						if grid[board[i+k][j+l]] > 1 {
							return false
						}
					}
				}
			}
		}
	}

	return true
}

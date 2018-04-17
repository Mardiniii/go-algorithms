package main

// This functions needs to verify all the positions in the
// row, column or box unless the given position

func checkRow(grid [N][N]int, n int, row int, col int) bool {
	for j := 0; j < N; j++ {
		if j != col && grid[row][j] == n {
			return false
		}
	}
	return true
}

func checkColumn(grid [N][N]int, n int, row int, col int) bool {
	for i := 0; i < N; i++ {
		if i != row && grid[i][col] == n {
			return false
		}
	}
	return true
}

func checkBox(grid [N][N]int, n int, rowStart int, colStart int, row int, col int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (i+rowStart != row) &&
				(j+colStart != col) &&
				(grid[i+rowStart][j+colStart] == n) {
				return false
			}
		}
	}
	return true
}

// CheckBoard checs the given grid is a valid solution for the Sudoku given problem
func CheckBoard(grid [N][N]int) bool {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if !checkRow(grid, grid[i][j], i, j) ||
				!checkColumn(grid, grid[i][j], i, j) ||
				!checkBox(grid, grid[i][j], i-(i%3), j-(j%3), i, j) {
				return false
			}
		}
	}
	return true
}

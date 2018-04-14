package main

import (
	"fmt"
	"strconv"
)

// N represents the length for each row and column
const N = 9

// PENDING used as identifier for pending positions
const PENDING = 0

var count int

func pendingPositions(grid [N][N]int) (bool, int, int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == PENDING {
				return true, i, j
			}
		}
	}
	return false, -1, -1
}

// validInRow checks if the value is not repeated in the given row
func validInRow(grid [N][N]int, n int, row int) bool {
	for j := 0; j < N; j++ {
		if grid[row][j] == n {
			return false
		}
	}
	return true
}

// validInColumn checks if the value is not repeated in the given col
func validInColumn(grid [N][N]int, n int, col int) bool {
	for i := 0; i < N; i++ {
		if grid[i][col] == n {
			return false
		}
	}
	return true
}

// validInColumn checks if the value is not repeated in the 3x3 box where is
// contain the current position
func validInBox(grid [N][N]int, n int, rowStart int, colStart int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+rowStart][j+colStart] == n {
				return false
			}
		}
	}
	return true
}

func valid(grid [N][N]int, n int, row int, col int) bool {
	if validInRow(grid, n, row) &&
		validInColumn(grid, n, col) &&
		validInBox(grid, n, row-(row%3), col-(col%3)) {
		return true
	}

	return false
}

// Exported methods

// NewSudoku creates a new grid from string input
func NewSudoku(input string) [N][N]int {
	var grid [N][N]int

	for i := 0; i < N; i++ {
		row := input[0:N]
		input = input[N:]
		for j := 0; j < N; j++ {
			grid[i][j] = int(row[j]) - 48
		}
	}
	count = 0

	return grid
}

// Solve algorithm
func Solve(grid [N][N]int) (bool, [N][N]int) {
	ch := make(chan [N][N]int)

	pending, row, col := pendingPositions(grid)
	if !pending {
		fmt.Println("Solve calls:", count)
		return true, grid
	}

	// Let's try to find a solution for the pending position
	// in row and col. We will try with digits from 1 to 9
	for i := 1; i <= 9; i++ {
		if valid(grid, i, row, col) {
			grid[row][col] = i
			go recursiveSolution(grid, ch)
		}
	}

	select {
	case solution := <-ch:
		fmt.Println("Solve calls:", count)
		return true, solution
	}
}

func recursiveSolution(grid [N][N]int, ch chan [N][N]int) {
	count++
	pending, row, col := pendingPositions(grid)

	if !pending && CheckBoard(grid) {
		ch <- grid
		close(ch)
		return
	}

	// Let's try to find a solution for the pending position
	// in row and col. We will try with digits from 1 to 9
	for i := 1; i <= 9; i++ {
		if valid(grid, i, row, col) {
			// If is valid assign this value to the given position
			grid[row][col] = i

			// Recursive call to Solve method. If the Sudoku is done return true
			go recursiveSolution(grid, ch)
		}
	}
}

// Print the board with its current state
func Print(grid [N][N]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Println()
	}
}

// ToString returns one line string with the current sudoku result
func ToString(grid [N][N]int) string {
	var r string

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			r += strconv.Itoa(grid[i][j])
		}
	}

	return r
}

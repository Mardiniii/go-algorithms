package main

import (
	"fmt"
	"strconv"
)

// N represents the length for each row and column
const N = 9

// PENDING used as identifier for pending positions
const PENDING = 0

// Sudoku struct contains the grid
type Sudoku struct {
	grid [N][N]int
}

// Unexported methods
func (s *Sudoku) buildGrid(input string) {
	for i := 0; i < N; i++ {
		row := input[0:N]
		input = input[N:]
		for j := 0; j < N; j++ {
			s.grid[i][j] = int(row[j]) - 48
		}
	}
}

func (s *Sudoku) pendingPositions() (bool, int, int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if s.grid[i][j] == PENDING {
				return true, i, j
			}
		}
	}
	return false, -1, -1
}

// validInRow checks if the value is not repeated in the given row
func (s *Sudoku) validInRow(n int, row int) bool {
	for j := 0; j < N; j++ {
		if s.grid[row][j] == n {
			return false
		}
	}
	return true
}

// validInColumn checks if the value is not repeated in the given col
func (s *Sudoku) validInColumn(n int, col int) bool {
	for i := 0; i < N; i++ {
		if s.grid[i][col] == n {
			return false
		}
	}
	return true
}

// validInColumn checks if the value is not repeated in the 3x3 box where is
// contain the current position
func (s *Sudoku) validInBox(n int, rowStart int, colStart int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if s.grid[i+rowStart][j+colStart] == n {
				return false
			}
		}
	}
	return true
}

func (s *Sudoku) valid(n int, row int, col int) bool {
	if s.validInRow(n, row) &&
		s.validInColumn(n, col) &&
		s.validInBox(n, row-(row%3), col-(col%3)) {
		return true
	}

	return false
}

// Exported methods

// NewSudoku creates a Sudoku with grid from input string
func NewSudoku(input string) *Sudoku {
	s := new(Sudoku)
	s.buildGrid(input)

	return s
}

// Solve algorithm
func (s *Sudoku) Solve() bool {
	pending, row, col := s.pendingPositions()
	if !pending {
		return true
	}

	// Let's try to find a solution for the pending position
	// in row and col. We will try with digits from 1 to 9
	for i := 1; i <= 9; i++ {
		// Check is the current digit is valid for the given position
		if s.valid(i, row, col) {
			// If is valid assign this value to the given position
			s.grid[row][col] = i

			// Recursive call to Solve method. If the Sudoku is done return true
			if s.Solve() {
				return true
			}

			// If the current value does not fixed the Sudoku we undo the changes
			// and try with the next one
			s.grid[row][col] = PENDING
		}
	}

	// We don't find any solution for the Sudoku problem
	return false
}

// Print the board with its current state
func (s *Sudoku) Print() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%d ", s.grid[i][j])
		}
		fmt.Println()
	}
}

// ToString returns one line string with the current sudoku result
func (s *Sudoku) ToString() string {
	var r string

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			r += strconv.Itoa(s.grid[i][j])
		}
	}

	return r
}

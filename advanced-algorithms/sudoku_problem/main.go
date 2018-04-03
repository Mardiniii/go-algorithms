package main

import "fmt"

func main() {
	var input string

	fmt.Println("Plase enter a valid grid(81 characters): ")
	fmt.Scanf("%s", &input)

	s := NewSudoku(input)

	fmt.Println("Sudoku board loaded successfully")
	s.Print()

	fmt.Println()

	s.Solve()

	s.Print()
}

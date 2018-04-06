package main

import (
	"fmt"
	"math"
)

// N is the number of cities
const N = 4

// AllVisited is the mask used to represent when
// all the cities have been visited
const AllVisited = (1 << N) - 1

// memoTable store the results already computed
var memoTable [16][N]int

func buildMatrix() [][]int {
	var n int

	fmt.Printf("Number of cities: ")
	fmt.Scanf("%d", &n)

	fmt.Println("Cities", n)

	// allocate composed 2d array
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("Position[%d][%d]: ", i+1, i+1)
			fmt.Scanf("%d", &matrix[i][j])
		}
	}

	return matrix
}

func tsp(m [][]int, mask int, pos uint) int {
	if mask == AllVisited {
		return m[pos][0]
	}

	if memoTable[mask][pos] != -1 {
		return memoTable[mask][pos]
	}
	var city uint
	var newCost int
	minCost := math.MaxInt64 // Greatest possible value

	// Let's go to an unvisited city
	for city = 0; city < N; city++ {
		cityAsBinary := (1 << city)

		if (mask & cityAsBinary) == 0 {
			newCost = m[pos][city] + tsp(m, mask|cityAsBinary, city)
			minCost = min(minCost, newCost)
			memoTable[mask][pos] = minCost
		}

	}

	return minCost
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// m := buildMatrix()
	// m := [][]int{
	// 	[]int{0, 20, 42, 25},
	// 	[]int{20, 0, 30, 34},
	// 	[]int{42, 30, 0, 10},
	// 	[]int{25, 34, 10, 0},
	// }

	m := [][]int{
		[]int{0, 10, 15, 20},
		[]int{10, 0, 35, 25},
		[]int{15, 35, 0, 30},
		[]int{20, 25, 30, 0},
	}

	for i := 0; i < (1 << N); i++ {
		for j := 0; j < N; j++ {
			memoTable[i][j] = -1
		}
	}

	cost := tsp(m, 1, 0)

	fmt.Println("The minimum cost is:", cost)
}

package main

import (
	"fmt"
)

type hanoiTower struct{}

var start []int
var aux []int
var end []int
var c int

func (ht *hanoiTower) Play(n int) {
	for i := 1; i <= n; i++ {
		start = append(start, i)
	}
	c = 0

	move(len(start), &start, &aux, &end)
}

// Move n elements from the start array to the end array
// using aux as axuliary column
func move(n int, start *[]int, aux *[]int, end *[]int) {
	if n < 1 {
		return
	}

	if n == 1 {
		e := (*start)[0]
		*start = (*start)[1:]
		*end = append([]int{e}, (*end)...)
		c++
		printMove()
	} else {
		move(n-1, start, end, aux)
		move(1, start, aux, end)
		move(n-1, aux, start, end)
	}
}

func printMove() {
	fmt.Println()
	fmt.Println("Move:", c)
	fmt.Println("Start:", start)
	fmt.Println("Aux:", aux)
	fmt.Println("End:", end)
}

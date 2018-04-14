package main

import (
	"fmt"
	"testing"
	"time"
)

func TestSudoku(t *testing.T) {
	testCases := []struct {
		problem  string
		solution string
	}{
		{
			"200370009009200007001004002050000800008000900006000040900100500800007600400089001",
			"284375169639218457571964382152496873348752916796831245967143528813527694425689731",
		},
		{
			"306508400520000000087000031003010080900863005050090600130000250000000074005206300",
			"316578492529134768487629531263415987974863125851792643138947256692351874745286319",
		},
		{
			"003080001020001000100300800200009070006240000050000100300900008000056090040000700",
			"693782541825461937174395862238519674916247385457638129361974258782156493549823716",
		},
	}

	for _, testCase := range testCases {
		grid := NewSudoku(testCase.problem)
		startTime := time.Now()
		_, solution := Solve(grid)
		endTime := time.Now()
		diff := endTime.Sub(startTime)
		fmt.Println("total time taken ", diff.Seconds(), "seconds")

		algorithmResult := ToString(solution)

		if testCase.solution != algorithmResult {
			t.Errorf("Sudoku is not solved, got: %v, want: %v.", algorithmResult, testCase.solution)
		}
	}
}

package main

import "testing"

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
	}

	for _, testCase := range testCases {
		s := NewSudoku(testCase.problem)
		s.Solve()

		algorithmResult := s.ToString()

		if testCase.solution != algorithmResult {
			t.Errorf("Sudoku is not solved, got: %v, want: %v.", algorithmResult, testCase.solution)
		}
	}
}

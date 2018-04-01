package fibonacci

import "testing"

func TestMemoizedFibonacci(t *testing.T) {
	testCases := []struct {
		index  int
		result int
	}{
		{0, 0},
		{1, 1},
		{8, 21},
		{9, 34},
		{10, 55},
		{11, 89},
		{37, 24157817},
	}

	for _, testCase := range testCases {
		algorithmResult := MemoizedFibonacci(testCase.index)

		if testCase.result != algorithmResult {
			t.Errorf(
				"Fibonnaci algorithm failed expected: %v, got: %v",
				testCase.result,
				algorithmResult)
		}
	}
}

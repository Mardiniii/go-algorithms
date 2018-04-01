package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		sortedList []int
		lookFor    int
		index      int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 5, 4},
		{[]int{0, 6, 7, 8}, 3, -1},
		{[]int{0, 1}, 1, 1},
		{[]int{1}, 3, -1},
		{[]int{}, 3, -1},
	}

	for _, testCase := range testCases {
		algorithmResult := BinarySearch(testCase.sortedList, testCase.lookFor)

		if testCase.index != algorithmResult {
			t.Errorf(
				"Binary search algorithm failed looking for: %v, in: %v, expected: %v, got: %v",
				testCase.lookFor, testCase.sortedList, testCase.index, algorithmResult)
		}
	}
}

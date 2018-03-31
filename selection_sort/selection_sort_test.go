package selectionsort

import "testing"

func TestSelectionSort(t *testing.T) {
	testCases := []struct {
		unsortedList []int
		sortedList   []int
	}{
		{[]int{3, 1, 4, 5, 2}, []int{1, 2, 3, 4, 5}},
		{[]int{6, 7, 8, 0}, []int{0, 6, 7, 8}},
		{[]int{1, 0}, []int{0, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, testCase := range testCases {
		algorithmResult := SelectionSort(testCase.unsortedList)

		for i, v := range testCase.sortedList {
			if v != algorithmResult[i] {
				t.Errorf("List is not ordered, got: %v, want: %v.", algorithmResult, testCase.sortedList)
			}
		}
	}
}

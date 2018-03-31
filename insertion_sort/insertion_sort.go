package insertionsort

// Input: Unsorted list
// Output: Sorted list
// How it works: Insertion sort is a simple sorting algorithm
// where each element from the list is compared with the other
// ones at the left to check if it is in the correct position.
// This algorithm is similar to order cards in our hands in real life.
// Time Complexity: O(n^2). It’s a very inefficient algorithm especially
// for lists with a big size.
// Space complexity: 0(1) since all the operations are done in place

// InsertionSort algorithm applied to the given unsorted list
func InsertionSort(list []int) []int {
	if len(list) == 0 {
		return []int{}
	}

	for i := 1; i < len(list); i++ {
		for j := i; j > 0; j-- {
			if list[j] < list[j-1] {
				lowerValue := list[j]
				list[j] = list[j-1]
				list[j-1] = lowerValue
			} else {
				break
			}
		}
	}

	return list
}

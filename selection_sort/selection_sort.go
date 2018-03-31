package selectionsort

// Input: Unsorted list
// Output: Sorted list
// How it works: Selection sort is one of the simplest algorithms
// to sort a list, this algorithm works by scanning a list of items
// for the smallest element, and then swapping that element for the
// one in first position. This continues with the remaining items
// until the list is sorted.
// Time Complexity: The complexity of selection sort ranges from O(1)
// when the list is already sorted to O(n^2) when the list is presorted
// in the reverse order you want. The O(1) complexity makes this an
// interesting choice if there’s a chance of sorting a pre-sorted
// list(or a list of equal objects) – which happens fairly often.
// Space complexity: 0(1) since all the operations are done in place
// on the given list

// SelectionSort algorithm applied to the given unsorted list
func SelectionSort(list []int) []int {
	for i := 0; i < len(list); i++ {
		minIndex := i
		for j := minIndex + 1; j < len(list); j++ {
			if list[j] < list[minIndex] {
				minIndex = j
			}

		}
		if minIndex != i {
			oldMinValue := list[i]
			list[i] = list[minIndex]
			list[minIndex] = oldMinValue
		}
	}

	return list
}

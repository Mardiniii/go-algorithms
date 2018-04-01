package quicksort

// Input: Unsorted list
// Output: Sorted list
// How it works: Quicksort is a divide and conquer algorithm that
// uses a pivoting technique to break the main list into smaller lists.
// These smaller lists use the pivoting technique until they are sorted.
// To choose the pivot any of the next options can be used:

// - Always pick first element as pivot.
// - Always pick last element as pivot (implemented below)
// - Pick a random element as pivot.
// - Pick median as pivot.

// After choose the pivot in the current list the elements greater than
// the pivot goes to the right and its position is replaced with the left
// element in the pivot. At the end we should have elements greater than
// the pivot in the right and elements lower than the pivot in the left.
// Time complexity: O(n log n) or O(n*2). The complexity of quicksort,
// however, is not constant because the pivot is determined at random and
// the partitioning of the list can put the algorithm at a disadvantage.
// In the worst case, quicksort is O(n^2) when the pivot is the smallest or
// largest element in the list. In the best case it's O(n log n), like merge sort.
// Space complexity: O(log(n))

// QuickSort algorithm applied to the given unsorted list
func QuickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	// Partition lists
	var left []int
	var right []int
	// Choose the pivot index and value
	p := len(list) - 1
	v := list[p]

	// Remove the pivot
	list = append(list[:p])

	for i := 0; i < len(list); i++ {
		if list[i] < v {
			left = append(left, list[i])
		} else {
			right = append(right, list[i])
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, []int{v}...), right...)
}

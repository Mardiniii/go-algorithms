package heapsort

// Input: Unsorted list
// Output: Sorted list
// How it works: Heap sort is a bit like selection sort in
// that it moves unsorted data to a sorted “partition” selectively.
// The difference, however, is that it uses a heap to do so.
// A heap is a tree structure where parent nodes in one level
// are either greater than (max heap) or less than (min heap)
// child nodes in descendant levels. We receive the unsorted list
// then we have to create the heap and start to extract the element
// on the heap from the top to the bottom to find the final sorted list result
// Time complexity: Heap sort is O(n log n), however it has an advantage
// over quicksort of being O(n log n) in the worst case scenario,
// whereas quicksort in the worst case in O(n^2).
// Space complexity: O(1)

// heapify method builds a max-heap from the element in position i
func heapify(list []int, n int, i int) {
	// Guess current node as the largest value
	largest := i
	// Find left and right children
	left := 2*i + 1
	right := 2*i + 2

	// Compare current node value with the left child
	if left < n && list[left] > list[largest] {
		largest = left
	}

	// Compare current node value with the right child
	if right < n && list[right] > list[largest] {
		largest = right
	}

	// Check if a swap is needed
	if largest != i {
		temp := list[i]
		list[i] = list[largest]
		list[largest] = temp

		// Recursive call to heapify the affected sub-tree
		heapify(list, n, largest)
	}
}

// HeapSort algorithm applied to the given unsorted list
func HeapSort(list []int) []int {
	// Build max heap from the last non leaf node index (n/2) - 1
	for i := len(list)/2 - 1; i >= 0; i-- {
		heapify(list, len(list), i)
	}

	// Apply HeapSort algorithm changing first with the last node.
	// Repeat the process exluding the last node. It is already in
	// its final position
	for i := len(list) - 1; i >= 0; i-- {
		temp := list[i]
		list[i] = list[0]
		list[0] = temp

		heapify(list, i, 0)
	}
	return list
}

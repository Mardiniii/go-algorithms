package bubblesort

// Input: Unsorted list
// Output: Sorted list
// How it works: Bubble sort compare each pair of continuous
// elements in the array until we don’t find more changes to be made.
// Time complexity: O(n^2). It’s a very inefficient algorithm especially
// for arrays with a big size
// Space complexity: 0(1) since all the operations are done in place
// on the given list

// BubbleSort algorithm applied to the given unsorted list
func BubbleSort(list []int) []int {
	for range list {
		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				smallestValue := list[i+1]
				list[i+1] = list[i]
				list[i] = smallestValue
			}
		}
	}

	return list
}

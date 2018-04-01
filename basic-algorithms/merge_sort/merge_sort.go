package mergesort

// Input: Unsorted list
// Output: Sorted list
// How it works: Merge Sort works by splitting all the elements
// in a list down to smaller, two-element lists which can then be
// sorted easily in one pass. The final step is to recursively merge
// these smaller lists back into a larger list
// Time complexity: O(n log n). This is a very efficient algorithm
// that uses divide and conquer approach to split the list in smaller(two elements)
// lists using O(log n) complexity. These small lists should be ordered
// which implies a O(n) operation.
// Space complexity: O(n)

func merge(l []int, r []int) []int {
	result := []int{}

	for len(l) > 0 || len(r) > 0 {
		if len(l) > 0 && len(r) > 0 {
			if l[0] < r[0] {
				result = append(result, l[0])
				l = l[1:]
			} else {
				result = append(result, r[0])
				r = r[1:]
			}
		} else if len(l) > 0 {
			result = append(result, l[0])
			l = l[1:]
		} else {
			result = append(result, r[0])
			r = r[1:]
		}
	}

	return result
}

// MergeSort algorithm applied to the given unsorted list
func MergeSort(list []int) []int {
	switch len(list) {
	case 0:
		return list
	case 1:
		return list
	default:
		m := len(list) / 2
		l := list[:m]
		r := list[m:]
		return merge(MergeSort(l), MergeSort(r))
	}
}

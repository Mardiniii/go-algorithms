package binarysearch

// BinarySearch algorithm applied to the given sorted list
// BinarySearch returns the index of the lookFor value if it
// is present in the list, otherwise it will return -1
func BinarySearch(l []int, lookFor int) int {
	switch len(l) {
	case 0:
		return -1
	case 1:
		if l[0] == lookFor {
			return 0
		}

		return -1
	default:
		min, max := 0, len(l)

		for min <= max {
			middle := (min + max) / 2

			if l[middle] == lookFor {
				return middle
			}

			if l[middle] < lookFor {
				min = middle + 1
			} else {
				max = middle - 1
			}
		}
	}

	return -1
}

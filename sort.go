package iterago

// Sort allow you to sort your slice with Merge Sort algorithm.
//
// Predicate MUST be comparison between the first and the second value.
//   - The comparison MUST check if the first value is bigger than the second value for ASC.
//   - The comparison MUST check if the first value is smaller than the second value for DESC.
//
// Currently doesn't support multithreading
func Sort[T any](values []T, predicate func(T, T) bool) []T {
	if len(values) <= 1 {
		return values
	}

	left := Sort(values[:len(values)/2], predicate)
	right := Sort(values[len(values)/2:], predicate)

	return mergeSort(left, right, predicate)
}

func mergeSort[T any](left []T, right []T, predicate func(T, T) bool) (sorted []T) {
	for len(left) > 0 && len(right) > 0 {
		if predicate(left[0], right[0]) {
			sorted = append(sorted, right[:1]...)
			right = right[1:]
			continue
		}

		sorted = append(sorted, left[:1]...)
		left = left[1:]
	}

	if len(left) == 0 {
		sorted = append(sorted, right...)
	}

	if len(right) == 0 {
		sorted = append(sorted, left...)
	}

	return
}

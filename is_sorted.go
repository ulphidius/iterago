package iterago

// IsSorted allow you to check if the slice is sorted.
//
// Predicate MUST be comparison between the first and the second value.
//   - The comparison MUST check if the first value is smaller or equal than the second value for ASC.
//   - The comparison MUST check if the first value is bigger or equal than the second value for DESC.
func IsSorted[T any](values []T, predicate func(T, T) bool) bool {
	if len(values) <= 1 {
		return true
	}

	return predicate(values[0], values[1]) && IsSorted(values[1:], predicate)
}

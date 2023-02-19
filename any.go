package iterago

// Any allow to check if all values of the slice aren't valid
func Any[T any](values []T, predicate func(T) bool) bool {
	if len(values) == 0 {
		return false
	}

	return anyHelper(values, predicate)
}

func anyHelper[T any](values []T, predicate func(T) bool) bool {
	if len(values) == 0 {
		return true
	}

	if !predicate(values[0]) {
		return true && anyHelper(values[1:], predicate)
	}

	return false
}

package iterago

// All allow to check if all values of the slice are valid
func All[T any](values []T, predicate func(T) bool) bool {
	if len(values) == 0 {
		return false
	}

	return allHelper(values, predicate)
}

func allHelper[T any](values []T, predicate func(T) bool) bool {
	if len(values) == 0 {
		return true
	}

	if predicate(values[0]) {
		return true && allHelper(values[1:], predicate)
	}

	return false
}

package iterago

// Partition allow you by using a predicate function to return one list with all valid values and another one with the rest.
func Partition[T any](values []T, predicate func(T) bool) (validated []T, invalidated []T) {
	for _, value := range values {
		if predicate(value) {
			validated = append(validated, value)
			continue
		}

		invalidated = append(invalidated, value)
	}

	return
}

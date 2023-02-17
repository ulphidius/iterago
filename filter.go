package iterago

func Filter[T any](values []T, predicate func(T) bool) []T {
	if len(values) == 0 {
		return nil
	}

	if predicate(values[0]) {
		return append(values[:1], Filter(values[1:], predicate)...)
	}

	return Filter(values[1:], predicate)
}

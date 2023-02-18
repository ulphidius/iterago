package iterago

func Partition[T any](values []T, predicate func(T) bool) ([]T, []T) {
	if len(values) == 0 {
		return nil, nil
	}

	validated, invalidated := Partition(values[1:], predicate)

	if predicate(values[0]) {
		return append(values[:1], validated...), invalidated
	}

	return validated, append(values[:1], invalidated...)
}

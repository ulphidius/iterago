package iterago

func Reduce[T any](values []T, accumulator T, predicate func(T, T) T) T {
	if len(values) == 0 {
		return accumulator
	}

	return Reduce(values[1:], predicate(accumulator, values[0]), predicate)
}

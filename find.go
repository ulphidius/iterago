package iterago

func Find[T any](values []T, predicate func(T) bool) Option[T] {
	if len(values) == 0 {
		return NewNoneOption[T]()
	}

	if predicate(values[0]) {
		return NewOption(values[0])
	}

	return Find(values[1:], predicate)
}

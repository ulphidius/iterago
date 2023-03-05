package iterago

func Foreach[T any](values []T, predicate func(T)) {
	if len(values) == 0 {
		return
	}

	predicate(values[0])

	Foreach(values[1:], predicate)
}

package iterago

func Map[T, G any](values []T, predicate func(T) G) []G {
	if len(values) == 0 {
		return nil
	}

	return append([]G{predicate(values[0])}, Map(values[1:], predicate)...)
}

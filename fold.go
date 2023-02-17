package iterago

func Fold[T, G any](values []T, accumulator G, predicate func(G, T) G) G {
	if len(values) == 0 {
		return accumulator
	}

	return Fold(values[1:], predicate(accumulator, values[0]), predicate)
}

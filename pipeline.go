package iterago

func FilterMap[T, G any](values []T, filterPredicate func(T) bool, mapPredicate func(T) G) []G {
	if len(values) == 0 {
		return nil
	}

	if filterPredicate(values[0]) {
		return append([]G{mapPredicate(values[0])}, FilterMap(values[1:], filterPredicate, mapPredicate)...)
	}

	return FilterMap(values[1:], filterPredicate, mapPredicate)
}

func FilterReduce[T any](values []T, accumulator T, filterPredicate func(T) bool, reducePredicate func(T, T) T) T {
	if len(values) == 0 {
		return accumulator
	}

	if filterPredicate(values[0]) {
		return FilterReduce(values[1:], reducePredicate(accumulator, values[0]), filterPredicate, reducePredicate)
	}

	return FilterReduce(values[1:], accumulator, filterPredicate, reducePredicate)
}

func FilterFold[T, G any](values []T, accumulator G, filterPredicate func(T) bool, foldPredicate func(G, T) G) G {
	if len(values) == 0 {
		return accumulator
	}

	if filterPredicate(values[0]) {
		return FilterFold(values[1:], foldPredicate(accumulator, values[0]), filterPredicate, foldPredicate)
	}

	return FilterFold(values[1:], accumulator, filterPredicate, foldPredicate)
}

func MapReduce[T, G any](values []T, accumulator G, mapPredicate func(T) G, reducePredicate func(G, G) G) G {
	if len(values) == 0 {
		return accumulator
	}

	return MapReduce(values[1:], reducePredicate(accumulator, mapPredicate(values[0])), mapPredicate, reducePredicate)
}

func MapFold[T, G, V any](values []T, accumulator V, mapPredicate func(T) G, foldPredicate func(V, G) V) V {
	if len(values) == 0 {
		return accumulator
	}

	return MapFold(values[1:], foldPredicate(accumulator, mapPredicate(values[0])), mapPredicate, foldPredicate)
}

package iterago

type FilterMapPredicates[T, G any] struct {
	Filter func(T) bool
	Map    func(T) G
}

type FilterReducePredicates[T any] struct {
	Filter func(T) bool
	Reduce func(T, T) T
}

type FilterFoldPredicates[T, G any] struct {
	Filter func(T) bool
	Fold   func(G, T) G
}

type MapReducePredicates[T, G any] struct {
	Map    func(T) G
	Reduce func(G, G) G
}

type MapFoldPrediactes[T, G, V any] struct {
	Map  func(T) G
	Fold func(V, G) V
}

type PartitionForeachPredicates[T any] struct {
	Filter      func(T) bool
	Validate    func(T) // Foreach predicate for valid values
	Invalidates func(T) // Foreach predicate for invalid values
}

func FilterMap[T, G any](values []T, predicates FilterMapPredicates[T, G]) []G {
	if len(values) == 0 {
		return nil
	}

	if predicates.Filter(values[0]) {
		return append([]G{predicates.Map(values[0])}, FilterMap(values[1:], predicates)...)
	}

	return FilterMap(values[1:], predicates)
}

func FilterReduce[T any](values []T, accumulator T, predicates FilterReducePredicates[T]) T {
	if len(values) == 0 {
		return accumulator
	}

	if predicates.Filter(values[0]) {
		return FilterReduce(values[1:], predicates.Reduce(accumulator, values[0]), predicates)
	}

	return FilterReduce(values[1:], accumulator, predicates)
}

func FilterFold[T, G any](values []T, accumulator G, predicates FilterFoldPredicates[T, G]) G {
	if len(values) == 0 {
		return accumulator
	}

	if predicates.Filter(values[0]) {
		return FilterFold(values[1:], predicates.Fold(accumulator, values[0]), predicates)
	}

	return FilterFold(values[1:], accumulator, predicates)
}

func MapReduce[T, G any](values []T, accumulator G, predicates MapReducePredicates[T, G]) G {
	if len(values) == 0 {
		return accumulator
	}

	return MapReduce(values[1:], predicates.Reduce(accumulator, predicates.Map(values[0])), predicates)
}

func PartitionForeach[T any](values []T, predicates PartitionForeachPredicates[T]) {
	if len(values) == 0 {
		return
	}

	if predicates.Filter(values[0]) {
		predicates.Validate(values[0])

		PartitionForeach(values[1:], predicates)
		return
	}

	predicates.Invalidates(values[0])
	PartitionForeach(values[1:], predicates)
}

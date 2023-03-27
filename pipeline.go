package iterago

import "sync"

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

type UniqueMergePredicates[T any, G Comparable] struct {
	Identifier func(T) G
	Merge      func(T, T) T
}

func FilterMap[T, G any](values []T, predicates FilterMapPredicates[T, G]) []G {
	if len(values) == 0 {
		return nil
	}

	if IteragoThreads > 1 {
		return filterMapMultithreads(IteragoThreads, values, predicates)
	}

	return filterMap(values, predicates)
}

func filterMapMultithreads[T, G any](threads uint, values []T, predicates FilterMapPredicates[T, G]) []G {
	size := getChunkSize(values, threads)
	chunks := split(values, size)

	var wg sync.WaitGroup
	var mx sync.Mutex
	var result []G
	for _, chunk := range chunks {
		wg.Add(1)

		go func(c []T, pre FilterMapPredicates[T, G]) {
			defer wg.Done()
			tmp := filterMap(c, pre)
			mx.Lock()
			result = append(result, tmp...)
			mx.Unlock()
		}(chunk, predicates)
	}
	wg.Wait()

	return result
}

func filterMap[T, G any](values []T, predicates FilterMapPredicates[T, G]) []G {
	if len(values) == 0 {
		return nil
	}

	if predicates.Filter(values[0]) {
		return append([]G{predicates.Map(values[0])}, filterMap(values[1:], predicates)...)
	}

	return filterMap(values[1:], predicates)
}

func FilterReduce[T any](values []T, accumulator T, predicates FilterReducePredicates[T]) T {
	if len(values) == 0 {
		return accumulator
	}

	if IteragoThreads > 1 {
		return filterReduceMultithreads(IteragoThreads, values, accumulator, predicates)
	}

	return *filterReduce(values, NewOption(accumulator), predicates)
}

func filterReduceMultithreads[T any](threads uint, values []T, accumulator T, predicates FilterReducePredicates[T]) T {
	size := getChunkSize(values, threads)
	chunks := split(values, size)

	var wg sync.WaitGroup
	var mx sync.Mutex
	var result T
	for index, chunk := range chunks {
		wg.Add(1)

		go func(i int, c []T, pre FilterReducePredicates[T]) {
			defer wg.Done()
			var tmp *T
			if i == 0 {
				tmp = filterReduce(c, NewOption(accumulator), pre)
			} else {
				tmp = filterReduce(c, NewNoneOption[T](), pre)
			}

			if tmp == nil {
				return
			}

			mx.Lock()
			result = pre.Reduce(result, *tmp)
			mx.Unlock()
		}(index, chunk, predicates)
	}
	wg.Wait()

	return result
}

func filterReduce[T any](values []T, accumulator Option[T], predicates FilterReducePredicates[T]) *T {
	if len(values) == 0 {
		if accumulator.IsNone() {
			return nil
		}

		tmp := accumulator.Unwrap()
		return &tmp
	}

	if predicates.Filter(values[0]) {
		if accumulator.IsNone() {
			return filterReduce(values[1:], NewOption(values[0]), predicates)
		}

		return filterReduce(values[1:], NewOption(predicates.Reduce(accumulator.Unwrap(), values[0])), predicates)
	}

	return filterReduce(values[1:], accumulator, predicates)
}

// FilterFold currently doesn't support multithreading
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

	if IteragoThreads > 1 {
		return mapReduceMultithreads(IteragoThreads, values, accumulator, predicates)
	}

	return mapReduce(values, NewOption(accumulator), predicates)
}

func mapReduceMultithreads[T, G any](threads uint, values []T, accumulator G, predicates MapReducePredicates[T, G]) G {
	size := getChunkSize(values, threads)
	chunks := split(values, size)

	var wg sync.WaitGroup
	var mx sync.Mutex
	var result G
	for index, chunk := range chunks {
		wg.Add(1)

		go func(i int, c []T, pre MapReducePredicates[T, G]) {
			defer wg.Done()
			var tmp G
			if i == 0 {
				tmp = mapReduce(c, NewOption(accumulator), pre)
			} else {
				tmp = mapReduce(c, NewNoneOption[G](), pre)
			}
			mx.Lock()
			result = pre.Reduce(result, tmp)
			mx.Unlock()
		}(index, chunk, predicates)
	}
	wg.Wait()

	return result
}

func mapReduce[T, G any](values []T, accumulator Option[G], predicates MapReducePredicates[T, G]) G {
	if len(values) == 0 {
		return accumulator.Unwrap()
	}

	if accumulator.IsNone() {
		return mapReduce(values[1:], NewOption(predicates.Map(values[0])), predicates)
	}

	return mapReduce(values[1:], NewOption(predicates.Reduce(accumulator.Unwrap(), predicates.Map(values[0]))), predicates)
}

func PartitionForeach[T any](values []T, predicates PartitionForeachPredicates[T]) {
	if len(values) == 0 {
		return
	}

	if IteragoThreads > 1 {
		partitionForeachMultithreads(IteragoThreads, values, predicates)
		return
	}

	partitionForeach(values, predicates)
}

func partitionForeachMultithreads[T any](threads uint, values []T, predicates PartitionForeachPredicates[T]) {
	size := getChunkSize(values, threads)
	chunks := split(values, size)

	var wg sync.WaitGroup
	for _, chunk := range chunks {
		wg.Add(1)

		go func(c []T, pre PartitionForeachPredicates[T]) {
			defer wg.Done()
			partitionForeach(c, pre)

		}(chunk, predicates)
	}
	wg.Wait()
}

func partitionForeach[T any](values []T, predicates PartitionForeachPredicates[T]) {
	if len(values) == 0 {
		return
	}

	if predicates.Filter(values[0]) {
		predicates.Validate(values[0])

		partitionForeach(values[1:], predicates)
		return
	}

	predicates.Invalidates(values[0])
	partitionForeach(values[1:], predicates)
}

// UniqueMerge currently doesn't support multithreading
func UniqueMerge[T any, G Comparable](values []T, predicates UniqueMergePredicates[T, G]) []T {
	if len(values) == 0 {
		return nil
	}

	mapper := uniqueMergeHelper(values, predicates, map[G]T{})
	result := []T{}
	for _, value := range mapper {
		result = append(result, value)
	}

	return result
}

func uniqueMergeHelper[T any, G Comparable](values []T, predicates UniqueMergePredicates[T, G], mapper map[G]T) map[G]T {
	if len(values) == 0 {
		return mapper
	}

	key := predicates.Identifier(values[0])
	if _, ok := mapper[key]; ok {
		mapper[key] = predicates.Merge(mapper[key], values[0])
		return uniqueMergeHelper(values[1:], predicates, mapper)
	}

	mapper[key] = values[0]
	return uniqueMergeHelper(values[1:], predicates, mapper)
}

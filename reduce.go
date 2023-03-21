package iterago

import "sync"

func Reduce[T any](values []T, accumulator T, predicate func(T, T) T) T {
	if len(values) == 0 {
		return accumulator
	}

	if IteragoThreads > 1 {
		return reduceMultithreads(IteragoThreads, values, accumulator, predicate)
	}

	return reduce(values, NewOption(accumulator), predicate)
}

func reduceMultithreads[T any](threads uint, values []T, accumulator T, predicate func(T, T) T) T {
	size := getChunkSize(values, threads)
	chunks := split(values, size)

	var wg sync.WaitGroup
	var mx sync.Mutex
	var result T
	for index, chunk := range chunks {
		wg.Add(1)

		go func(i int, c []T, pre func(T, T) T) {
			defer wg.Done()
			var tmp T
			if i == 0 {
				tmp = reduce(c, NewOption(accumulator), pre)
			} else {
				tmp = reduce(c, NewNoneOption[T](), pre)
			}
			mx.Lock()
			result = pre(result, tmp)
			mx.Unlock()
		}(index, chunk, predicate)
	}
	wg.Wait()

	return result
}

func reduce[T any](values []T, accumulator Option[T], predicate func(T, T) T) T {
	if len(values) == 0 {
		return accumulator.Unwrap()
	}

	if accumulator.IsNone() {
		return reduce(values[1:], NewOption(values[0]), predicate)
	}

	return reduce(values[1:], NewOption(predicate(accumulator.Unwrap(), values[0])), predicate)
}

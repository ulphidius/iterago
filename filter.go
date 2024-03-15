package iterago

import (
	"sync"
)

func Filter[T any](values []T, predicate func(T) bool) []T {
	if len(values) == 0 {
		return nil
	}

	if IteragoThreads > 1 {
		return filterMultithreads(IteragoThreads, values, predicate)
	}

	return filter(values, predicate)
}

func filterMultithreads[T any](threads uint, values []T, predicate func(T) bool) []T {
	size := getChunkSize(values, threads)
	chunks := split(values, size)

	var wg sync.WaitGroup
	var mx sync.Mutex
	result := []T{}
	for _, chunk := range chunks {
		wg.Add(1)

		go func(c []T, pre func(T) bool) {
			defer wg.Done()
			tmp := filter(c, pre)
			mx.Lock()
			result = append(result, tmp...)
			mx.Unlock()
		}(chunk, predicate)
	}
	wg.Wait()

	return result
}

func filter[T any](values []T, predicate func(T) bool) []T {
	if len(values) == 0 {
		return nil
	}

	if predicate(values[0]) {
		return append([]T{values[0]}, filter(values[1:], predicate)...)
	}

	return filter(values[1:], predicate)
}

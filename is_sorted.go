package iterago

import "sync"

// IsSorted allow you to check if the slice is sorted.
//
// Predicate MUST be comparison between the first and the second value.
//   - The comparison MUST check if the first value is smaller or equal than the second value for ASC.
//   - The comparison MUST check if the first value is bigger or equal than the second value for DESC.
func IsSorted[T any](values []T, predicate func(T, T) bool) bool {
	if len(values) <= 1 {
		return true
	}

	if IteragoThreads > 1 {
		return isSortedMultithreads(IteragoThreads, values, predicate)
	}

	return isSorted(values, predicate)
}

func isSortedMultithreads[T any](threads uint, values []T, predicate func(T, T) bool) bool {
	size := getChunkSize(values, threads)
	chunks := split(values, size)

	var wg sync.WaitGroup
	var mx sync.Mutex
	result := true
	for _, chunk := range chunks {
		wg.Add(1)

		go func(c []T, pre func(T, T) bool) {
			defer wg.Done()
			tmp := isSorted(c, pre)
			mx.Lock()
			result = result && tmp
			mx.Unlock()
		}(chunk, predicate)
	}
	wg.Wait()

	return result
}

func isSorted[T any](values []T, predicate func(T, T) bool) bool {
	if len(values) <= 1 {
		return true
	}

	return predicate(values[0], values[1]) && isSorted(values[1:], predicate)
}

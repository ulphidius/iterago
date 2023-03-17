package iterago

import "sync"

// Any allow to check if all values of the slice aren't valid
func Any[T any](values []T, predicate func(T) bool) bool {
	if len(values) == 0 {
		return false
	}

	if iteragoThreads > 1 {
		return allMultithreads(iteragoThreads, values, predicate)
	}

	return anyHelper(values, predicate)
}

func anyMultithreads[T any](threads uint, values []T, predicate func(T) bool) bool {
	size := getChunkSize(values, threads)
	chunks := split(values, size)

	var wg sync.WaitGroup
	var mx sync.Mutex
	result := true
	for _, chunk := range chunks {
		wg.Add(1)

		go func(c []T, pre func(T) bool) {
			defer wg.Done()
			tmp := anyHelper(c, pre)
			mx.Lock()
			result = result && tmp
			mx.Unlock()
		}(chunk, predicate)
	}
	wg.Wait()

	return result
}

func anyHelper[T any](values []T, predicate func(T) bool) bool {
	if len(values) == 0 {
		return true
	}

	if !predicate(values[0]) {
		return true && anyHelper(values[1:], predicate)
	}

	return false
}

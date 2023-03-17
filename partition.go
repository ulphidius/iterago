package iterago

import "sync"

// Partition allow you by using a predicate function to return one list with all valid values and another one with the rest.
func Partition[T any](values []T, predicate func(T) bool) (validated []T, invalidated []T) {
	if len(values) == 0 {
		return nil, nil
	}

	if iteragoThreads > 1 {
		return partitionMultithreads(iteragoThreads, values, predicate)
	}

	return partition(values, predicate)
}

func partitionMultithreads[T any](threads uint, values []T, predicate func(T) bool) (validated []T, invalidated []T) {
	size := getChunkSize(values, threads)
	chunks := split(values, size)

	var wg sync.WaitGroup
	var mx sync.Mutex
	for _, chunk := range chunks {
		wg.Add(1)

		go func(c []T, pre func(T) bool) {
			defer wg.Done()
			valids, invalids := partition(c, pre)
			mx.Lock()
			validated = append(validated, valids...)
			invalidated = append(invalidated, invalids...)
			mx.Unlock()
		}(chunk, predicate)
	}
	wg.Wait()

	return
}

func partition[T any](values []T, predicate func(T) bool) (validated []T, invalidated []T) {
	for _, value := range values {
		if predicate(value) {
			validated = append(validated, value)
			continue
		}

		invalidated = append(invalidated, value)
	}

	return
}

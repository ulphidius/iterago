package iterago

import "sync"

func Map[T, G any](values []T, predicate func(T) G) []G {
	if len(values) == 0 {
		return nil
	}

	if iteragoThreads > 1 {
		return mapperMultithreads(iteragoThreads, values, predicate)
	}

	return mapper(values, predicate)
}

func mapperMultithreads[T, G any](threads uint, values []T, predicate func(T) G) []G {
	size := getChunkSize(values, threads)
	chunks := split(values, size)

	var wg sync.WaitGroup
	var mx sync.Mutex
	result := []G{}
	for _, chunk := range chunks {
		wg.Add(1)

		go func(c []T, pre func(T) G) {
			defer wg.Done()
			tmp := mapper(c, pre)
			mx.Lock()
			result = append(result, tmp...)
			mx.Unlock()
		}(chunk, predicate)
	}
	wg.Wait()

	return result
}

func mapper[T, G any](values []T, predicate func(T) G) []G {
	if len(values) == 0 {
		return nil
	}

	return append([]G{predicate(values[0])}, Map(values[1:], predicate)...)
}

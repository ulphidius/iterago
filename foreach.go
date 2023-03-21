package iterago

import "sync"

func Foreach[T any](values []T, predicate func(T)) {
	if len(values) == 0 {
		return
	}

	if IteragoThreads > 1 {
		foreachMultithreads(IteragoThreads, values, predicate)
		return
	}

	foreach(values, predicate)
}

func foreachMultithreads[T any](threads uint, values []T, predicate func(T)) {
	size := getChunkSize(values, threads)
	chunks := split(values, size)

	var wg sync.WaitGroup
	for _, chunk := range chunks {
		wg.Add(1)

		go func(c []T, pre func(T)) {
			defer wg.Done()
			foreach(c, pre)
		}(chunk, predicate)
	}
	wg.Wait()
}

func foreach[T any](values []T, predicate func(T)) {
	if len(values) == 0 {
		return
	}

	predicate(values[0])

	foreach(values[1:], predicate)
}

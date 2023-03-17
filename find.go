package iterago

import "sync"

// Find allow you to return the first that matches the predicate condition if it exists
func Find[T any](values []T, predicate func(T) bool) Option[T] {
	if len(values) == 0 {
		return NewNoneOption[T]()
	}

	if iteragoThreads > 1 {
		return findMultithreads(iteragoThreads, values, predicate)
	}

	return find(values, predicate)
}

func findMultithreads[T any](threads uint, values []T, predicate func(T) bool) Option[T] {
	size := getChunkSize(values, threads)
	chunks := split(values, size)

	var wg sync.WaitGroup
	var mx sync.Mutex
	result := NewNoneOption[T]()
	for index, chunk := range chunks {
		wg.Add(1)

		go func(i int, c []T, pre func(T) bool) {
			defer wg.Done()
			tmp := find(c, pre)
			mx.Lock()
			if tmp.IsSome() {
				result = tmp
			}
			mx.Unlock()
		}(index, chunk, predicate)
	}
	wg.Wait()

	return result
}

func find[T any](values []T, predicate func(T) bool) Option[T] {
	if len(values) == 0 {
		return NewNoneOption[T]()
	}

	if predicate(values[0]) {
		return NewOption(values[0])
	}

	return Find(values[1:], predicate)
}

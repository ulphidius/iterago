package iterago

import "sync"

// Zip allow you to merge to array into an array of Pair
func Zip[T any](first []T, second []T) []Pair[Option[T]] {
	if len(first) == 0 && len(second) == 0 {
		return nil
	}

	if iteragoThreads > 1 {
		return zipMultithreads(iteragoThreads, first, second)
	}

	return zip(first, second)
}

func zipMultithreads[T any](threads uint, first []T, second []T) (result []Pair[Option[T]]) {
	firstChunks, secondChunks := zipGetChunks(first, second, threads)
	maxLenght := func(f int, s int) int {
		if f > s {
			return f
		}

		return s
	}(len(firstChunks), len(secondChunks))

	var wg sync.WaitGroup
	var mx sync.Mutex
	for i := 0; i < maxLenght; i += 1 {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()
			var tmp []Pair[Option[T]] = nil
			if index > len(firstChunks)-1 {
				tmp = zip(nil, secondChunks[index])
			}
			if index > len(secondChunks)-1 {
				tmp = zip(firstChunks[index], nil)
			}
			if tmp == nil {
				tmp = zip(firstChunks[index], secondChunks[index])
			}
			mx.Lock()
			result = append(result, tmp...)
			mx.Unlock()
		}(i)
	}
	wg.Wait()

	return
}

func zipGetChunks[T any](first, second []T, threads uint) ([][]T, [][]T) {
	if len(first) == 0 {
		secondSize := getChunkSize(second, threads)
		return nil, split(second, secondSize)
	}

	if len(second) == 0 {
		firstSize := getChunkSize(first, threads)
		return split(first, firstSize), nil
	}

	firtSize := getChunkSize(first, threads)
	firstChunks := split(first, firtSize)
	secondSize := len(firstChunks[0])
	secondChunks := chunks(second, uint(secondSize))

	return firstChunks, secondChunks
}

func zip[T any](first []T, second []T) []Pair[Option[T]] {
	if len(first) == 0 && len(second) == 0 {
		return nil
	}

	if len(first) == 0 {
		return append(
			[]Pair[Option[T]]{
				NewPair(
					NewNoneOption[T](),
					NewOption(second[0]),
				),
			},
			zip(nil, second[1:])...,
		)
	}

	if len(second) == 0 {
		return append(
			[]Pair[Option[T]]{
				NewPair(
					NewOption(first[0]),
					NewNoneOption[T](),
				),
			},
			zip(first[1:], nil)...,
		)
	}

	return append(
		[]Pair[Option[T]]{
			NewPair(
				NewOption(first[0]),
				NewOption(second[0]),
			),
		},
		zip(first[1:], second[1:])...,
	)
}

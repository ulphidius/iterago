package iterago

import "sync"

// Zip allow you to merge to array into an array of Pair
func Zip[T, G any](first []T, second []G) []Pair[Option[T], Option[G]] {
	if len(first) == 0 && len(second) == 0 {
		return nil
	}

	if IteragoThreads > 1 {
		return zipMultithreads(IteragoThreads, first, second)
	}

	return zip(first, second)
}

func zipMultithreads[T, G any](threads uint, first []T, second []G) (result []Pair[Option[T], Option[G]]) {
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
			var tmp []Pair[Option[T], Option[G]] = nil
			if index > len(firstChunks)-1 {
				tmp = zip[T, G](nil, secondChunks[index])
			}
			if index > len(secondChunks)-1 {
				tmp = zip[T, G](firstChunks[index], nil)
			}
			if tmp == nil {
				tmp = zip[T, G](firstChunks[index], secondChunks[index])
			}
			mx.Lock()
			result = append(result, tmp...)
			mx.Unlock()
		}(i)
	}
	wg.Wait()

	return
}

func zipGetChunks[T, G any](first []T, second []G, threads uint) ([][]T, [][]G) {
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

func zip[T, G any](first []T, second []G) []Pair[Option[T], Option[G]] {
	if len(first) == 0 && len(second) == 0 {
		return nil
	}

	if len(first) == 0 {
		return append(
			[]Pair[Option[T], Option[G]]{
				NewPair(
					NewNoneOption[T](),
					NewOption(second[0]),
				),
			},
			zip[T, G](nil, second[1:])...,
		)
	}

	if len(second) == 0 {
		return append(
			[]Pair[Option[T], Option[G]]{
				NewPair(
					NewOption(first[0]),
					NewNoneOption[G](),
				),
			},
			zip[T, G](first[1:], nil)...,
		)
	}

	return append(
		[]Pair[Option[T], Option[G]]{
			NewPair(
				NewOption(first[0]),
				NewOption(second[0]),
			),
		},
		zip(first[1:], second[1:])...,
	)
}

func MapIntoZip[T Comparable, G any](m map[T]G) []Pair[Option[T], Option[G]] {
	keys := []T{}
	values := []G{}

	for key, value := range m {
		keys = append(keys, key)
		values = append(values, value)
	}

	return Zip(keys, values)
}

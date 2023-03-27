package iterago

// Unique remove all duplicated values by taking the first one
//
// currently doesn't support multithreading
func Unique[T any, G Comparable](values []T, predicate func(T) G) []T {
	if len(values) == 0 {
		return nil
	}

	mapper := uniqueHelper(values, predicate, map[G]T{})
	result := []T{}
	for _, value := range mapper {
		result = append(result, value)
	}

	return result
}

func uniqueHelper[T any, G Comparable](values []T, predicate func(T) G, mapper map[G]T) map[G]T {
	if len(values) == 0 {
		return mapper
	}

	key := predicate(values[0])
	if _, ok := mapper[key]; ok {
		return uniqueHelper(values[1:], predicate, mapper)
	}

	mapper[key] = values[0]
	return uniqueHelper(values[1:], predicate, mapper)
}

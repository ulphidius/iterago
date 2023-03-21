package iterago

// Enumerate currently doesn't support multithreading
func Enumerate[T any](values []T) []EnumPair[T] {
	return enumerateHelper(values, 0)
}

func enumerateHelper[T any](values []T, iteration uint) []EnumPair[T] {
	if len(values) == 0 {
		return nil
	}

	return append(
		[]EnumPair[T]{NewEnumPair(iteration, values[0])},
		enumerateHelper(values[1:], iteration+1)...,
	)
}

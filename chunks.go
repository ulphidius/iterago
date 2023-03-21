package iterago

// Chunks allow you to split your array into an array of sub arrays of a specific size
//
// Currently doesn't support multithreading
func Chunks[T any](values []T, size uint) [][]T {
	if len(values) == 0 {
		return nil
	}

	if size == 0 || uint(len(values)) < size {
		return [][]T{values}
	}

	return chunks(values, size)
}

func chunks[T any](values []T, size uint) [][]T {
	if len(values) == 0 {
		return nil
	}

	if size == 0 || uint(len(values)) < size {
		return [][]T{values}
	}

	return append([][]T{
		values[:size],
	}, chunks(values[size:], size)...)
}

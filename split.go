package iterago

func Split[T any](values []T, number uint) [][]T {
	if len(values) == 0 {
		return nil
	}

	if number <= 1 {
		return [][]T{values}
	}

	size := getChunkSize(values, number)

	return split(values, size)
}

func split[T any](values []T, size int) [][]T {
	if len(values) == 0 {
		return nil
	}

	if size > len(values) {
		return [][]T{values}
	}

	return append([][]T{values[:size]}, split(values[size:], size)...)
}

func getChunkSize[T any](values []T, number uint) int {
	if len(values) == 0 {
		return 0
	}

	numberOfElements := len(values) / int(number)

	if len(values)%int(number) != 0 {
		return numberOfElements + 1
	}

	return numberOfElements
}

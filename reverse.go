package iterago

// Currently doesn't support mutithreading
func Reverse[T any](values []T) []T {
	if len(values) == 0 {
		return nil
	}

	if len(values) == 1 {
		return values
	}

	return append([]T{values[len(values)-1]}, Reverse(values[:len(values)-1])...)
}

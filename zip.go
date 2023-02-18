package iterago

// Zip allow you to merge to array into an array of Pair
func Zip[T any](first []T, second []T) []Pair[Option[T]] {
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
			Zip(nil, second[1:])...,
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
			Zip(first[1:], nil)...,
		)
	}

	return append(
		[]Pair[Option[T]]{
			NewPair(
				NewOption(first[0]),
				NewOption(second[0]),
			),
		},
		Zip(first[1:], second[1:])...,
	)
}

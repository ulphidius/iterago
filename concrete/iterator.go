package concrete

import "github.com/ulphidius/iterago/interfaces"

type Iterator[T any] struct {
	current T
	next    interfaces.Option[*Iterator[T]]
}

func (iter *Iterator[T]) Next() interfaces.Option[*Iterator[T]] {
	if iter == nil || !iter.HasNext() {
		return interfaces.NewNoneOption[*Iterator[T]]()
	}

	next, _ := iter.next.Unwrap()

	return interfaces.NewOption(next)
}

func (iter *Iterator[T]) HasNext() bool {
	if iter == nil {
		return false
	}

	return iter.next.IsSome()
}

func (iter *Iterator[T]) Filter(predicate func(T) bool) interfaces.Option[*Filtered[T]] {
	if iter == nil {
		return interfaces.NewNoneOption[*Filtered[T]]()
	}

	if !iter.HasNext() {
		return interfaces.NewOption(
			&Filtered[T]{
				current:    iter.current,
				next:       interfaces.NewNoneOption[*Filtered[T]](),
				predicates: []func(T) bool{predicate},
			},
		)
	}
	next, _ := iter.Next().Unwrap()

	return interfaces.NewOption(
		&Filtered[T]{
			current:    iter.current,
			next:       next.Filter(predicate),
			predicates: []func(T) bool{predicate},
		},
	)
}

func SliceUintIntoIter(values []uint) interfaces.Option[*Iterator[uint]] {
	if len(values) == 0 {
		return interfaces.NewNoneOption[*Iterator[uint]]()
	}

	return interfaces.NewOption(
		&Iterator[uint]{
			current: values[0],
			next:    SliceUintIntoIter(values[1:]),
		},
	)
}

package concrete

import (
	"github.com/ulphidius/iterago/interfaces"
)

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
			NewFilteredItem(
				iter.current,
				interfaces.NewNoneOption[*Filtered[T]](),
				predicate,
			),
		)
	}
	next, _ := iter.Next().Unwrap()

	return interfaces.NewOption(
		NewFilteredItem(
			iter.current,
			next.Filter(predicate),
			predicate,
		),
	)
}

func (iter *Iterator[T]) Map(predicate func(T) interface{}) interfaces.Option[*Mapper[T, interface{}]] {
}

func (iter *Iterator[T]) Collect() []T {
	if iter == nil {
		return nil
	}

	if iter.HasNext() {
		next, _ := iter.Next().Unwrap()
		return append([]T{iter.current}, next.Collect()...)
	}

	return []T{iter.current}
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

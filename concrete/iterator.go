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

func (iter *Iterator[T]) Map(predicate func(T) any) interfaces.Option[*Mapper[T, any]] {
	if iter == nil {
		return interfaces.NewNoneOption[*Mapper[T, any]]()
	}

	if !iter.HasNext() {
		return interfaces.NewOption(
			NewMapperItem(
				iter.current,
				interfaces.NewNoneOption[*Mapper[T, any]](),
				predicate,
			),
		)
	}

	next, _ := iter.Next().Unwrap()

	return interfaces.NewOption(
		NewMapperItem(
			iter.current,
			next.Map(predicate),
			predicate,
		),
	)
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

func (iter *Iterator[T]) Reduce(accumulator T, predicate func(x T, y T) T) T {
	if iter == nil {
		return accumulator
	}

	if iter.HasNext() {
		next, _ := iter.Next().Unwrap()
		return next.Reduce(predicate(accumulator, iter.current), predicate)
	}

	return predicate(accumulator, iter.current)
}

func SliceIntoIter[T any](values []T) interfaces.Option[*Iterator[T]] {
	if len(values) == 0 {
		return interfaces.NewNoneOption[*Iterator[T]]()
	}

	return interfaces.NewOption(
		&Iterator[T]{
			current: values[0],
			next:    SliceIntoIter(values[1:]),
		},
	)
}

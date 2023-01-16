package concrete

import "github.com/ulphidius/iterago"

type Iterator[T any] struct {
	current T
	next    iterago.Option[*Iterator[T]]
}

func (iter *Iterator[T]) Next() iterago.Option[*Iterator[T]] {
	if iter.HasNext() {
		return iter.next
	}

	return iterago.NewNoneOption[*Iterator[T]]()
}

func (iter *Iterator[T]) HasNext() bool {
	if iter == nil {
		return false
	}

	return iter.next.IsSome()
}

func SliceUintIntoIter(values []uint) iterago.Option[*Iterator[uint]] {
	if len(values) == 0 {
		return iterago.NewNoneOption[*Iterator[uint]]()
	}

	return iterago.NewOption(
		&Iterator[uint]{
			current: values[0],
			next:    SliceUintIntoIter(values[1:]),
		},
	)
}

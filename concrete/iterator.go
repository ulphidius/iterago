package concrete

import "github.com/ulphidius/iterago/interfaces"

type Iterator[T any] struct {
	current T
	next    interfaces.Option[*Iterator[T]]
}

func (iter *Iterator[T]) Next() interfaces.Option[*Iterator[T]] {
	if iter.HasNext() {
		return iter.next
	}

	return interfaces.NewNoneOption[*Iterator[T]]()
}

func (iter *Iterator[T]) HasNext() bool {
	if iter == nil {
		return false
	}

	return iter.next.IsSome()
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

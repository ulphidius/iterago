package concrete

import (
	"github.com/ulphidius/iterago/interfaces"
)

type Filtered[T any] struct {
	current    T
	next       interfaces.Option[*Filtered[T]]
	predicates []func(T) bool
	validated  bool
}

func (iter *Filtered[T]) Next() interfaces.Option[*Filtered[T]] {
	if iter == nil || !iter.HasNext() {
		return interfaces.NewNoneOption[*Filtered[T]]()
	}

	next, _ := iter.next.Unwrap()

	for _, predicate := range next.predicates {
		if predicate(next.current) {
			next.validated = true
		}
	}

	return interfaces.NewOption(next)
}

func (iter *Filtered[T]) HasNext() bool {
	if iter == nil {
		return false
	}

	return iter.next.IsSome()
}

func (iter *Filtered[T]) equal(value *Filtered[T]) bool {
	if iter == nil && value == nil {
		return true
	}

	next := (iter.HasNext() && value.HasNext()) || (!iter.HasNext() && !value.HasNext())

	predicateResult := false

	predicateResult = len(iter.predicates) == len(value.predicates)

	if predicateResult {
		for index := range iter.predicates {
			predicateResult = iter.predicates[index](iter.current) == value.predicates[index](value.current)
		}
	}

	return compare(value.current, iter.current) &&
		value.validated == iter.validated &&
		next &&
		predicateResult

}

func compare(a interface{}, b interface{}) bool {
	return a == b
}

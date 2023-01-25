package concrete

import "github.com/ulphidius/iterago/interfaces"

type Transform[T any, G any] func(T) G

type Mapper[T any, G any] struct {
	current   T
	next      interfaces.Option[*Mapper[T, G]]
	predicate func(T) G
	transform G
}

func NewMapperItem[T any, G any](
	value T,
	next interfaces.Option[*Mapper[T, G]],
	predicate func(T) G,
) *Mapper[T, G] {
	return &Mapper[T, G]{
		current:   value,
		next:      next,
		predicate: predicate,
		transform: predicate(value),
	}
}

func (iter *Mapper[T, G]) Next() interfaces.Option[*Mapper[T, G]] {
	if iter == nil || !iter.HasNext() {
		return interfaces.NewNoneOption[*Mapper[T, G]]()
	}

	next, _ := iter.next.Unwrap()
	iter.transform = iter.predicate(next.current)

	return interfaces.NewOption(next)
}

func (iter *Mapper[T, G]) HasNext() bool {
	if iter == nil {
		return false
	}

	return iter.next.IsSome()
}

func (iter *Mapper[T, G]) Filter(predicate func(G) bool) interfaces.Option[*Filtered[G]] {
	if iter == nil {
		return interfaces.NewNoneOption[*Filtered[G]]()
	}

	if !iter.HasNext() {
		return interfaces.NewOption(
			NewFilteredItem(
				iter.transform,
				interfaces.NewNoneOption[*Filtered[G]](),
				predicate,
			),
		)
	}

	next, _ := iter.Next().Unwrap()

	return interfaces.NewOption(
		NewFilteredItem(
			iter.transform,
			next.Filter(predicate),
			predicate,
		),
	)
}

func (iter *Mapper[T, G]) Collect() []G {
	if iter == nil {
		return nil
	}

	if iter.HasNext() {
		next, _ := iter.Next().Unwrap()
		return append([]G{iter.transform}, next.Collect()...)
	}

	return []G{iter.transform}
}

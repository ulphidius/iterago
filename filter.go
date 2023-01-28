package iterago

type Filter[T any] struct {
	current   T
	next      Option[*Filter[T]]
	predicate func(T) bool
	validated bool
}

func NewFilterItem[T any](
	value T,
	next Option[*Filter[T]],
	predicate func(x T) bool,
) *Filter[T] {
	return &Filter[T]{
		current:   value,
		next:      next,
		predicate: predicate,
		validated: predicate(value),
	}
}

func (iter *Filter[T]) Next() Option[*Filter[T]] {
	if iter == nil || !iter.HasNext() {
		return NewNoneOption[*Filter[T]]()
	}

	next, _ := iter.next.Unwrap()

	if next.predicate(next.current) {
		next.validated = true
	}

	return NewOption(next)
}

func (iter *Filter[T]) HasNext() bool {
	if iter == nil {
		return false
	}

	return iter.next.IsSome()
}

func (iter *Filter[T]) Filter(predicate func(T) bool) *Filter[T] {
	if iter == nil {
		return nil
	}

	if iter.HasNext() {
		next, _ := iter.Next().Unwrap()
		filtered := next.Filter(predicate)
		wrapped := func() Option[*Filter[T]] {
			if filtered == nil {
				return NewNoneOption[*Filter[T]]()
			}

			return NewOption(filtered)
		}()

		if iter.validated {
			return NewFilterItem(
				iter.current,
				wrapped,
				predicate,
			)
		}

		return next.Filter(predicate)
	}

	if iter.validated {
		return NewFilterItem(
			iter.current,
			NewNoneOption[*Filter[T]](),
			predicate,
		)
	}

	return nil
}

func (iter *Filter[T]) Map(predicate func(T) any) *Mapper[T, any] {
	if iter == nil {
		return nil
	}

	if !iter.HasNext() {
		if iter.validated {
			return NewMapperItem(
				iter.current,
				NewNoneOption[*Mapper[T, any]](),
				predicate,
			)
		}

		return nil
	}

	next, _ := iter.Next().Unwrap()
	filtered := next.Map(predicate)
	wrapped := func() Option[*Mapper[T, any]] {
		if filtered == nil {
			return NewNoneOption[*Mapper[T, any]]()
		}

		return NewOption(filtered)
	}()

	if iter.validated {
		return NewMapperItem(
			iter.current,
			wrapped,
			predicate,
		)
	}

	return next.Map(predicate)
}

func (iter *Filter[T]) Collect() []T {
	if iter == nil {
		return nil
	}

	if iter.HasNext() {
		next, _ := iter.Next().Unwrap()
		if iter.validated {
			return append([]T{iter.current}, next.Collect()...)
		}
		return next.Collect()
	}

	if iter.validated {
		return []T{iter.current}
	}

	return nil
}

func (iter *Filter[T]) Reduce(accumulator T, predicate func(x T, y T) T) T {
	if iter == nil {
		return accumulator
	}

	if iter.HasNext() {
		next, _ := iter.Next().Unwrap()
		if iter.validated {
			return next.Reduce(predicate(accumulator, iter.current), predicate)
		}
		return next.Reduce(accumulator, predicate)
	}

	return predicate(accumulator, iter.current)
}

func (iter *Filter[T]) equal(value *Filter[T]) bool {
	if iter == nil && value == nil {
		return true
	}

	next := (iter.HasNext() && value.HasNext()) || (!iter.HasNext() && !value.HasNext())

	return compare(value.current, iter.current) &&
		value.validated == iter.validated &&
		next &&
		iter.predicate(iter.current) == value.predicate(value.current)
}

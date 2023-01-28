package iterago

// Concrete implementation of a basic Iterator.
//
// It takes a value of a any Type and has a Option of a another Iter.
type Iter[T any] struct {
	current T
	next    Option[*Iter[T]]
}

// Explores Iterator by returning the next value
func (iter *Iter[T]) Next() Option[*Iter[T]] {
	if iter == nil || !iter.HasNext() {
		return NewNoneOption[*Iter[T]]()
	}

	next, _ := iter.next.Unwrap()

	return NewOption(next)
}

// Checks if the next value is setted.
func (iter *Iter[T]) HasNext() bool {
	if iter == nil {
		return false
	}

	return iter.next.IsSome()
}

// Creates an Iterator which validate the values.
//
// This function use a lambda function to validate the Iterator value.
func (iter *Iter[T]) Filter(predicate func(T) bool) *Filter[T] {
	if iter == nil {
		return nil
	}

	if !iter.HasNext() {
		return NewFilterItem(
			iter.current,
			NewNoneOption[*Filter[T]](),
			predicate,
		)
	}

	next, _ := iter.Next().Unwrap()
	filtered := next.Filter(predicate)
	wrapped := func() Option[*Filter[T]] {
		if filtered == nil {
			return NewNoneOption[*Filter[T]]()
		}

		return NewOption(filtered)
	}()

	return NewFilterItem(
		iter.current,
		wrapped,
		predicate,
	)
}

// Creates an Iterator which transform the values.
//
// This function use a lambda function to transform the Iterator value into another one.
func (iter *Iter[T]) Map(predicate func(T) any) *Mapper[T, any] {
	if iter == nil {
		return nil
	}

	if !iter.HasNext() {
		return NewMapperItem(
			iter.current,
			NewNoneOption[*Mapper[T, any]](),
			predicate,
		)
	}

	next, _ := iter.Next().Unwrap()
	filtered := next.Map(predicate)
	wrapped := func() Option[*Mapper[T, any]] {
		if filtered == nil {
			return NewNoneOption[*Mapper[T, any]]()
		}

		return NewOption(filtered)
	}()

	return NewMapperItem(
		iter.current,
		wrapped,
		predicate,
	)
}

// Consumes an Iterator into a Slice of values.
func (iter *Iter[T]) Collect() []T {
	if iter == nil {
		return nil
	}

	if iter.HasNext() {
		next, _ := iter.Next().Unwrap()
		return append([]T{iter.current}, next.Collect()...)
	}

	return []T{iter.current}
}

// Consumes an Iterator into a value.
func (iter *Iter[T]) Reduce(accumulator T, predicate func(x T, y T) T) T {
	if iter == nil {
		return accumulator
	}

	if iter.HasNext() {
		next, _ := iter.Next().Unwrap()
		return next.Reduce(predicate(accumulator, iter.current), predicate)
	}

	return predicate(accumulator, iter.current)
}

// Convert a Slice of generic values into an Iterator.
func SliceIntoIter[T any](values []T) Option[*Iter[T]] {
	if len(values) == 0 {
		return NewNoneOption[*Iter[T]]()
	}

	return NewOption(
		&Iter[T]{
			current: values[0],
			next:    SliceIntoIter(values[1:]),
		},
	)
}

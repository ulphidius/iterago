package iterago

import (
	"fmt"
)

// Concrete implementation of a basic Iterator.
//
// It takes a value of a any Type and has a Option of a another Iter.
// It takes a predication function which transforms the value into another one.
type Mapper[T any, G any] struct {
	current   T
	next      Option[*Mapper[T, G]]
	predicate func(T) G
	transform G
}

func NewMapperItem[T any, G any](
	value T,
	next Option[*Mapper[T, G]],
	predicate func(T) G,
) *Mapper[T, G] {
	return &Mapper[T, G]{
		current:   value,
		next:      next,
		predicate: predicate,
		transform: predicate(value),
	}
}

// Explores Iterator by returning the next value.
//
// It transforms the next value with the predicate function.
func (iter *Mapper[T, G]) Next() Option[*Mapper[T, G]] {
	if iter == nil || !iter.HasNext() {
		return NewNoneOption[*Mapper[T, G]]()
	}

	next, _ := iter.next.Unwrap()
	next.transform = iter.predicate(next.current)

	return NewOption(next)
}

func (iter *Mapper[T, G]) compute() Option[*Mapper[T, G]] {
	if iter == nil {
		return NewNoneOption[*Mapper[T, G]]()
	}

	iter.transform = iter.predicate(iter.current)

	return NewOption(iter)
}

// Checks if the next value is setted.
func (iter *Mapper[T, G]) HasNext() bool {
	if iter == nil {
		return false
	}

	return iter.next.IsSome()
}

// Creates an Iterator which validate the values.
//
// This function use a lambda function to validate the Iterator value.
// Only the transformed values are converted into Filter.
func (iter *Mapper[T, G]) Filter(predicate func(G) bool) *Filter[G] {
	if iter == nil {
		return nil
	}

	if !iter.HasNext() {
		return NewFilterItem(
			iter.transform,
			NewNoneOption[*Filter[G]](),
			predicate,
		)
	}

	next, _ := iter.Next().Unwrap()
	filtered := next.Filter(predicate)
	wrapped := func() Option[*Filter[G]] {
		if filtered == nil {
			return NewNoneOption[*Filter[G]]()
		}

		return NewOption(filtered)
	}()

	return NewFilterItem(
		iter.transform,
		wrapped,
		predicate,
	)
}

// Creates an Iterator which transform the values.
//
// This function use a lambda function to transform the Iterator value into another one.
// Only the transformed values are converted into Mapper.
func (iter *Mapper[T, G]) Map(predicate func(G) any) *Mapper[G, any] {
	if iter == nil {
		return nil
	}

	current, _ := iter.compute().Unwrap()

	if !current.HasNext() {
		return NewMapperItem(
			current.transform,
			NewNoneOption[*Mapper[G, any]](),
			predicate,
		)
	}

	next, _ := current.Next().Unwrap()
	filtered := next.Map(predicate)
	wrapped := func() Option[*Mapper[G, any]] {
		if filtered == nil {
			return NewNoneOption[*Mapper[G, any]]()
		}

		return NewOption(filtered)
	}()

	return NewMapperItem(
		current.transform,
		wrapped,
		predicate,
	)
}

// Consumes an Iterator into a Slice of values.
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

// Consumes an Iterator into a value.
func (iter *Mapper[T, G]) Reduce(accumulator G, predicate func(x G, y G) G) G {
	if iter == nil {
		return accumulator
	}

	if iter.HasNext() {
		next, _ := iter.Next().Unwrap()
		return next.Reduce(predicate(accumulator, iter.transform), predicate)
	}

	return predicate(accumulator, iter.transform)
}

func (iter *Mapper[T, G]) equal(value *Mapper[T, G]) bool {
	if iter == nil && value == nil {
		return true
	}

	next := (iter.HasNext() && value.HasNext()) || (!iter.HasNext() && !value.HasNext())

	predicateResult := fmt.Sprintf("%v", iter.predicate(iter.current)) == fmt.Sprintf("%v", value.predicate(value.current))

	return compare(value.current, iter.current) &&
		compare(value.transform, value.transform) &&
		next &&
		predicateResult

}

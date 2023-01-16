package interfaces

import "errors"

type Optional uint

const (
	None Optional = iota
	Some
)

type Option[T any] struct {
	Status Optional
	Value  T
}

func NewOption[T any](value T) Option[T] {
	return Option[T]{
		Status: Some,
		Value:  value,
	}
}

func NewNoneOption[T any]() Option[T] {
	return Option[T]{
		Status: None,
	}
}

func (opt Option[T]) IsSome() bool {
	return opt.Status == Some
}

func (opt Option[T]) IsNone() bool {
	return opt.Status == None
}

func (opt Option[T]) Unwrap() (T, error) {
	if opt.Status == None {
		return opt.Value, errors.New(ErrUnwrapNoneOption)
	}

	return opt.Value, nil
}

type Iterator[T any] interface {
	Next() Option[T]
	HasNext() bool
}

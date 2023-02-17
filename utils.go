package iterago

// Enumeration which defined the Option struct state.
type Optional uint

const (
	None Optional = iota // Represents the absence of value
	Some                 // Represents the presence of value
)

// Representation of an Optional value,
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

// Return the option value if option is some or return an error
func (opt Option[T]) Unwrap() T {
	if opt.Status == None {
		panic(ErrUnwrapNoneOption)
	}

	return opt.Value
}

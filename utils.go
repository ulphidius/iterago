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

// Return the option value if option is Some or panic if None
func (opt Option[T]) Unwrap() T {
	if opt.Status == None {
		panic(ErrUnwrapNoneOption)
	}

	return opt.Value
}

// Representation of 2 values
type Pair[T, G any] struct {
	First  T
	Second G
}

func NewPair[T, G any](first T, second G) Pair[T, G] {
	return Pair[T, G]{
		First:  first,
		Second: second,
	}
}

// Retrived Pair values as a tuple
func (pair Pair[T, G]) Unwrap() (T, G) {
	return pair.First, pair.Second
}

// Representation Pair value of an index and a value
type EnumPair[T any] struct {
	Index uint
	Value T
}

func NewEnumPair[T any](index uint, value T) EnumPair[T] {
	return EnumPair[T]{
		Index: index,
		Value: value,
	}
}

// MapIntoList convert a map into two list which contain the keys and the value
//
// currently doesn't support multithreading
func MapIntoList[T Comparable, G any](m map[T]G) ([]T, []G) {
	if len(m) == 0 {
		return nil, nil
	}

	keys := []T{}
	values := []G{}

	for key, value := range m {
		keys = append(keys, key)
		values = append(values, value)
	}

	return keys, values
}

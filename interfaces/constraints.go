package interfaces

type Unsigned interface {
	uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

type Signed interface {
	int | int8 | int16 | int32 | int64
}

type Float interface {
	float32 | float64
}

type Comparable interface {
	Unsigned | Signed | Float | string
}

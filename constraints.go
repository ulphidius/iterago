package iterago

type Unsigned interface {
	uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

type UnsignedSlice interface {
	[]uint | []uint8 | []uint16 | []uint32 | []uint64 | []uintptr
}

type Signed interface {
	int | int8 | int16 | int32 | int64
}

type SignedSlice interface {
	[]int | []int8 | []int16 | []int32 | []int64
}

type Float interface {
	float32 | float64
}

type FloatSlice interface {
	[]float32 | []float64
}

type Slice interface {
	UnsignedSlice | SignedSlice | FloatSlice | []string | []any
}

type Ordered interface {
	Unsigned | Signed | Float
}

type Comparable interface {
	Ordered | string
}

type Computable interface {
	Unsigned | Signed | Float
}

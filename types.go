package utils

// Numbers
type UInteger interface {
	uint8 | uint16 | uint32 | uint64
}

type Integer interface {
	int8 | int16 | int32 | int64
}

type Float interface {
	float32 | float64
}

type Number interface {
	UInteger | Integer | Float
}

// Primitives
type Primitive interface {
	string | Number | bool
}

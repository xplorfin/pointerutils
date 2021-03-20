package pointerutils

func FromInt(n int) *int {
	return &n
}

func ToInt(n *int) int {
	if n != nil {
		return *n
	}

	return 0
}

func FromInt8(n int8) *int8 {
	return &n
}

func ToInt8(n *int8) int8 {
	if n != nil {
		return *n
	}

	return 0
}

func FromInt16(n int16) *int16 {
	return &n
}

func ToInt16(n *int16) int16 {
	if n != nil {
		return *n
	}

	return 0
}

func FromInt32(n int32) *int32 {
	return &n
}

func ToInt32(n *int32) int32 {
	if n != nil {
		return *n
	}

	return 0
}

func FromInt64(n int64) *int64 {
	return &n
}

func ToInt64(n *int64) int64 {
	if n != nil {
		return *n
	}

	return 0
}

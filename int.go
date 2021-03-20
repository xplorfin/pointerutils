package pointerutils

// FromInt makes an int pointer from an int
func FromInt(n int) *int {
	return &n
}

// ToInt makes an int from an int pointer
func ToInt(n *int) int {
	if n != nil {
		return *n
	}

	return 0
}

// FromInt8 makes an int 8 pointer from an int 8
func FromInt8(n int8) *int8 {
	return &n
}

// ToInt8 makes an int 8 from an int 8 pointer
func ToInt8(n *int8) int8 {
	if n != nil {
		return *n
	}

	return 0
}

// FromInt16 makes an int 16 pointer from an int
func FromInt16(n int16) *int16 {
	return &n
}

// ToInt16 makes an int 16 from an int 16 pointer
func ToInt16(n *int16) int16 {
	if n != nil {
		return *n
	}

	return 0
}

// FromInt32 makes an int 32 pointer from an int
func FromInt32(n int32) *int32 {
	return &n
}

// ToInt32 makes an int 32 from an int 32 pointer
func ToInt32(n *int32) int32 {
	if n != nil {
		return *n
	}

	return 0
}

// FromInt64 makes an int 64 pointer from an int 64
func FromInt64(n int64) *int64 {
	return &n
}

// ToInt64 makes an int from an int 64 pointer
func ToInt64(n *int64) int64 {
	if n != nil {
		return *n
	}

	return 0
}

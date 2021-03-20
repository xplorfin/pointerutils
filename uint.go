package pointerutils

// FromUint makes a uint pointer from a uint
func FromUint(n uint) *uint {
	return &n
}

// ToUint makes a uint from a uint pointer
func ToUint(n *uint) uint {
	if n != nil {
		return *n
	}

	return 0
}

// FromUint8 makes a uint 8 pointer from a uint 8
func FromUint8(n uint8) *uint8 {
	return &n
}

// ToUint8 makes a uint 8 from a uint 8 pointer
func ToUint8(n *uint8) uint8 {
	if n != nil {
		return *n
	}

	return 0
}

// FromUint16 makes a uint 16 pointer from a uint 16
func FromUint16(n uint16) *uint16 {
	return &n
}

// ToUint16 makes a uint 16 from a uint 16 pointer
func ToUint16(n *uint16) uint16 {
	if n != nil {
		return *n
	}

	return 0
}

// FromUint32 makes a uint 32 pointer from a uint 32
func FromUint32(n uint32) *uint32 {
	return &n
}

// ToUint32 makes a uint 32 from a uint 32 pointer
func ToUint32(n *uint32) uint32 {
	if n != nil {
		return *n
	}

	return 0
}

// FromUint64 makes a uint 64 pointer from a uint 64
func FromUint64(n uint64) *uint64 {
	return &n
}

// ToUint64 makes a uint64 from a uint 64 pointer
func ToUint64(n *uint64) uint64 {
	if n != nil {
		return *n
	}

	return 0
}

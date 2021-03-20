package pointerutils

func FromUint(n uint) *uint {
	return &n
}

func ToUint(n *uint) uint {
	if n != nil {
		return *n
	}

	return 0
}

func FromUint8(n uint8) *uint8 {
	return &n
}

func ToUint8(n *uint8) uint8 {
	if n != nil {
		return *n
	}

	return 0
}

func FromUint16(n uint16) *uint16 {
	return &n
}

func ToUint16(n *uint16) uint16 {
	if n != nil {
		return *n
	}

	return 0
}

func FromUint32(n uint32) *uint32 {
	return &n
}

func ToUint32(n *uint32) uint32 {
	if n != nil {
		return *n
	}

	return 0
}

func FromUint64(n uint64) *uint64 {
	return &n
}

func ToUint64(n *uint64) uint64 {
	if n != nil {
		return *n
	}

	return 0
}

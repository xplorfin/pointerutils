package pointerutils

// FromBool makes a bool pointer from a bool
func FromBool(b bool) *bool {
	return &b
}

// ToBool makes a bool from a bool pointer
func ToBool(b *bool) bool {
	if b != nil {
		return *b
	}

	return false
}

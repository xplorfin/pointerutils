package pointerutils

func FromBool(b bool) *bool {
	return &b
}

func ToBool(b *bool) bool {
	if b != nil {
		return *b
	}

	return false
}

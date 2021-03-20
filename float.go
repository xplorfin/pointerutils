package pointerutils

// FromFloat32 makes a float 32 pointer from a float32
func FromFloat32(f float32) *float32 {
	return &f
}

// ToFloat32 makes a float 32 from a float 32 pointer
func ToFloat32(f *float32) float32 {
	if f != nil {
		return *f
	}

	return 0.
}

// FromFloat64 makes a float 64 from a float 64 pointer
func FromFloat64(f float64) *float64 {
	return &f
}

// ToFloat64 makes a float 64 from a float 64 pointer
func ToFloat64(f *float64) float64 {
	if f != nil {
		return *f
	}

	return 0.
}

package pointerutils

func FromFloat32(f float32) *float32 {
	return &f
}

func ToFloat32(f *float32) float32 {
	if f != nil {
		return *f
	}

	return 0.
}

func FromFloat64(f float64) *float64 {
	return &f
}

func ToFloat64(f *float64) float64 {
	if f != nil {
		return *f
	}

	return 0.
}

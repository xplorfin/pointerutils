package pointerutils

func FromString(s string) *string {
	return &s
}

func ToString(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}

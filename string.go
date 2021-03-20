package pointerutils

// FromString makes a string pointer from a string
func FromString(s string) *string {
	return &s
}

// ToString makes a string from a string pointer
func ToString(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}

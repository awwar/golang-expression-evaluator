package utility

func AsPtr[T any](v T) *T {
	return &v
}

package pointer

func Ref[T any](v T) *T {
	return &v
}

func Deref[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}

func Cast[T any](v any) *T {
	if v == nil {
		return nil
	}
	return v.(*T)
}

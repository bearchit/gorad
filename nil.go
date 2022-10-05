package gorad

func Ptr[T comparable](v T) *T {
	return &v
}

func Val[T comparable](p *T) T {
	if p == nil {
		var zv T
		return zv
	}
	return *p
}

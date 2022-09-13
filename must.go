package gorad

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

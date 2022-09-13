package gorad

import "context"

type Repository[T any] interface {
	Save(ctx context.Context, aggregate *T) error
	Find(ctx context.Context, id ID) (*T, error)
}

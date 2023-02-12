package must

import (
	"testing"
)

type that[T any] struct {
	m *Must
	v any
}

func That[T any](t testing.TB, value T) that[T] {
	return that[T]{New(t), value}
}

func (t that[T]) Equal(a any) that[T] {
	t.m.Equal(t.v, a)
	return t
}

func (t that[T]) NotEqual(a any) that[T] {
	t.m.NotEqual(t.v, a)
	return t
}

func (t that[T]) Nil() that[T] {
	t.m.Nil(t.v)
	return t
}

func (t that[T]) NotNil() that[T] {
	t.m.NotNil(t.v)
	return t
}

func (t that[T]) True() that[T] {
	t.m.Equal(t.v, true)
	return t
}

func (t that[T]) False() that[T] {
	t.m.Equal(t.v, false)
	return t
}

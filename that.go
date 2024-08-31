package should

import (
	"testing"
)

type that[T any] struct {
	m *Should
	v any
}

// That need testing.TB and value as parameters
// and "that" chainning instance will be returned
// with the value parameter as single assertable data
func That[T any](t testing.TB, value T) that[T] {
	return that[T]{New(t), value}
}

// Equal check are that's value and a equal, if not it will be failed
func (t that[T]) Equal(a any) that[T] {
	t.m.Equal(t.v, a)
	return t
}

// NotEqual reverse of Equal which is failed if that's value and a are equal
func (t that[T]) NotEqual(a any) that[T] {
	t.m.NotEqual(t.v, a)
	return t
}

// Nil check is that's value nil, if not it will be failed
func (t that[T]) Nil() that[T] {
	t.m.Nil(t.v)
	return t
}

// NotNil reverse of Nil which is failed if that's value nil
func (t that[T]) NotNil() that[T] {
	t.m.NotNil(t.v)
	return t
}

// True check is that's value true, if not it will be failed
func (t that[T]) True() that[T] {
	t.m.Equal(t.v, true)
	return t
}

// False check is that's value false, if not it will be failed
func (t that[T]) False() that[T] {
	t.m.Equal(t.v, false)
	return t
}

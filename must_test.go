package must

import "testing"

type mockT struct {
	testing.T
	failed bool
}

func Test_Must(t *testing.T) {
	t.Run("Equal", func(t *testing.T) {
		must := New(&mockT{})
		must.Equal(1, 1)
		New(t).False(must.t.Failed())

		must.Equal(1, 2)
		New(t).True(must.t.Failed())
	})

	t.Run("Not Equal", func(t *testing.T) {
		must := New(&mockT{})
		must.NotEqual(1, 2)
		New(t).False(must.t.Failed())

		must.NotEqual(1, 1)
		New(t).True(must.t.Failed())
	})

	t.Run("Nil", func(t *testing.T) {
		must := New(t)
		must.Nil(nil)

		var b []byte
		must.Nil(b)
	})

	t.Run("Not Nil", func(t *testing.T) {
		must := New(t)
		must.NotNil(1)
	})

	t.Run("True", func(t *testing.T) {
		must := New(t)
		must.True(true)
	})

	t.Run("False", func(t *testing.T) {
		must := New(t)
		must.False(false)
	})
}

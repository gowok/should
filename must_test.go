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
}

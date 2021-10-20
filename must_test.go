package must

import (
	"testing"
)

type mockT struct {
	testing.T
	failed bool
}

func (m mockT) logAndFail(args ...interface{}) {
	m.Fail()
}

func Test_Must(t *testing.T) {
	t.Run("callerinfo", func(t *testing.T) {
		must := New(t)
		_, _, ok := must.callerinfo()
		must.True(ok)
	})

	t.Run("Equal", func(t *testing.T) {
		must := New(&mockT{})
		must.Equal(1, 1)

		must = New(&mockT{})
		must.Equal(1, 2)

		must = New(&mockT{})
		must.Equal(&struct{}{}, &struct{}{})
	})

	t.Run("Not Equal", func(t *testing.T) {
		must := New(&mockT{})
		must.NotEqual(1, 2)

		must = New(&mockT{})
		must.NotEqual(1, 1)
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
		must.NotNil(&struct{}{})
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

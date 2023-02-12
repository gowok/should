package must

import (
	"testing"
)

type mockT struct {
	testing.T
	failed bool
}

func Test_Must(t *testing.T) {
	t.Run("callerinfo", func(t *testing.T) {
		must := New(t)
		_, _, ok := must.callerinfo()
		must.True(ok)
	})

	t.Run("logAndFail", func(t *testing.T) {
		tt := &mockT{}
		must := New(tt)
		must.logAndFail()
		False(t, tt.failed)
	})

	t.Run("Equal", func(t *testing.T) {
		tt := &mockT{}
		must := New(tt)
		must.Equal(1, 1)
		must.Equal(&struct{}{}, &struct{}{})
		False(t, tt.failed)
	})

	t.Run("Not Equal", func(t *testing.T) {
		tt := &mockT{}
		must := New(tt)
		must.NotEqual(1, 2)
		must.NotEqual(&struct{ i int }{1}, &struct{ i int }{2})
		False(t, tt.failed)
	})

	t.Run("Nil", func(t *testing.T) {
		tt := &mockT{}
		must := New(tt)
		must.Nil(nil)

		var b []byte
		must.Nil(b)
		False(t, tt.failed)
	})

	t.Run("Not Nil", func(t *testing.T) {
		tt := &mockT{}
		must := New(tt)
		must.NotNil(1)
		must.NotNil(&struct{}{})
		False(t, tt.failed)
	})

	t.Run("True", func(t *testing.T) {
		tt := &mockT{}
		must := New(t)
		must.True(true)
		False(t, tt.failed)
	})

	t.Run("False", func(t *testing.T) {
		tt := &mockT{}
		must := New(t)
		must.False(false)
		False(t, tt.failed)
	})
}

package should

import (
	"testing"
)

type mockT struct {
	testing.T
	failed bool
}

func Test_Should(t *testing.T) {
	t.Run("callerinfo", func(t *testing.T) {
		should := New(t)
		_, _, ok := should.callerinfo()
		should.True(ok)
	})

	t.Run("logAndFail", func(t *testing.T) {
		tt := &mockT{}
		should := New(tt)
		should.logAndFail()
		False(t, tt.failed)
	})

	t.Run("Equal", func(t *testing.T) {
		tt := &mockT{}
		should := New(tt)
		should.Equal(1, 1)
		should.Equal(&struct{}{}, &struct{}{})
		False(t, tt.failed)
	})

	t.Run("Not Equal", func(t *testing.T) {
		tt := &mockT{}
		should := New(tt)
		should.NotEqual(1, 2)
		should.NotEqual(&struct{ i int }{1}, &struct{ i int }{2})
		False(t, tt.failed)
	})

	t.Run("Nil", func(t *testing.T) {
		tt := &mockT{}
		should := New(tt)
		should.Nil(nil)

		var b []byte
		should.Nil(b)
		False(t, tt.failed)
	})

	t.Run("Not Nil", func(t *testing.T) {
		tt := &mockT{}
		should := New(tt)
		should.NotNil(1)
		should.NotNil(&struct{}{})
		False(t, tt.failed)
	})

	t.Run("True", func(t *testing.T) {
		tt := &mockT{}
		should := New(t)
		should.True(true)
		False(t, tt.failed)
	})

	t.Run("False", func(t *testing.T) {
		tt := &mockT{}
		should := New(t)
		should.False(false)
		False(t, tt.failed)
	})
}

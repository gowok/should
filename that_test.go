package should

import "testing"

func Test_That(t *testing.T) {
	tt := &mockT{}
	that := That(tt, 1)
	that.NotNil().Equal(1)

	that2 := That[*int](tt, nil)
	that2.Nil().NotEqual(1)

	that3 := That(tt, true)
	that3.True()

	that4 := That(tt, false)
	that4.False()
}

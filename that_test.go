package must

import "testing"

func Test_That(t *testing.T) {
	that := That(&mockT{}, 1)
	that.NotNil().Equal(1)

	that2 := That[*int](&mockT{}, nil)
	that2.Nil().NotEqual(1)
}

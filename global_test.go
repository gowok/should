package should

import "testing"

func Test_Global(t *testing.T) {
	t.Run("Equal", func(t *testing.T) {
		Equal(&mockT{}, 1, 1)
	})

	t.Run("Not Equal", func(t *testing.T) {
		NotEqual(&mockT{}, 1, 2)
	})

	t.Run("Nil", func(t *testing.T) {
		Nil(&mockT{}, nil)
	})

	t.Run("Not Nil", func(t *testing.T) {
		NotNil(&mockT{}, 1)
	})

	t.Run("True", func(t *testing.T) {
		True(&mockT{}, true)
	})

	t.Run("False", func(t *testing.T) {
		False(&mockT{}, false)
	})
}

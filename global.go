package should

import (
	"sync"
	"testing"
)

type shouldSync struct {
	m *Should
	l sync.RWMutex
}

func (m *shouldSync) setT(t testing.TB) {
	m.l.Lock()
	m.m.t = t
	m.l.Unlock()
}

var globalShould = shouldSync{m: &Should{}}

// Equal check are a and b equal, if not it will be failed
func Equal(t testing.TB, a, b any) {
	globalShould.setT(t)
	globalShould.m.Equal(a, b)
}

// NotEqual reverse of Equal which is failed if a and b are equal
func NotEqual(t testing.TB, a, b any) {
	globalShould.setT(t)
	globalShould.m.NotEqual(a, b)
}

// Nil check is a nil, if not it will be failed
func Nil(t testing.TB, a any) {
	globalShould.setT(t)
	globalShould.m.Nil(a)
}

// NotNil reverse of Nil which is failed if a nil
func NotNil(t testing.TB, a any) {
	globalShould.setT(t)
	globalShould.m.NotNil(a)
}

// True check is a true, if not it will be failed
func True(t testing.TB, a bool) {
	globalShould.setT(t)
	globalShould.m.True(a)
}

// False check is a false, if not it will be failed
func False(t testing.TB, a bool) {
	globalShould.setT(t)
	globalShould.m.False(a)
}

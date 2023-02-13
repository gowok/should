package must

import (
	"sync"
	"testing"
)

type mustSync struct {
	m *Must
	l sync.RWMutex
}

func (m *mustSync) setT(t testing.TB) {
	m.l.Lock()
	m.m.t = t
	m.l.Unlock()
}

var globalMust = mustSync{m: &Must{}}

// Equal check are a and b equal, if not it will be failed
func Equal(t testing.TB, a, b any) {
	globalMust.setT(t)
	globalMust.m.Equal(a, b)
}

// NotEqual reverse of Equal which is failed if a and b are equal
func NotEqual(t testing.TB, a, b any) {
	globalMust.setT(t)
	globalMust.m.NotEqual(a, b)
}

// Nil check is a nil, if not it will be failed
func Nil(t testing.TB, a any) {
	globalMust.setT(t)
	globalMust.m.Nil(a)
}

// NotNil reverse of Nil which is failed if a nil
func NotNil(t testing.TB, a any) {
	globalMust.setT(t)
	globalMust.m.NotNil(a)
}

// True check is a true, if not it will be failed
func True(t testing.TB, a bool) {
	globalMust.setT(t)
	globalMust.m.True(a)
}

// False check is a false, if not it will be failed
func False(t testing.TB, a bool) {
	globalMust.setT(t)
	globalMust.m.False(a)
}

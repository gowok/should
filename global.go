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

func Equal(t testing.TB, a, b any) {
	globalMust.setT(t)
	globalMust.m.Equal(a, b)
}

func NotEqual(t testing.TB, a, b any) {
	globalMust.setT(t)
	globalMust.m.NotEqual(a, b)
}

func Nil(t testing.TB, a any) {
	globalMust.setT(t)
	globalMust.m.Nil(a)
}

func NotNil(t testing.TB, a any) {
	globalMust.setT(t)
	globalMust.m.NotNil(a)
}

func True(t testing.TB, a bool) {
	globalMust.setT(t)
	globalMust.m.True(a)
}

func False(t testing.TB, a bool) {
	globalMust.setT(t)
	globalMust.m.False(a)
}

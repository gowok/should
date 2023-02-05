package must

import (
	"fmt"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"testing"
)

type Must struct {
	t testing.TB
}

func New(t testing.TB) *Must {
	return &Must{
		t: t,
	}
}

func isNil(a any) bool {
	if a == nil {
		return true
	}

	val := reflect.ValueOf(a)
	kind := val.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && val.IsNil() {
		return true
	}

	return false
}

func isEqual(a, b any) bool {
	if isNil(a) && isNil(b) {
		return true
	}

	if reflect.DeepEqual(a, b) {
		return true
	}

	return reflect.ValueOf(a) == reflect.ValueOf(b)
}

var reIsSourceFile = regexp.MustCompile("must\\.go$")

func (m Must) callerinfo() (path string, line int, ok bool) {
	for i := 0; ; i++ {
		_, path, line, ok = runtime.Caller(i)
		if !ok {
			return
		}

		if reIsSourceFile.MatchString(path) {
			continue
		}

		return path, line, true
	}
}

func (m Must) logAndFail(args ...any) {
	path, line, ok := m.callerinfo()
	if !ok {
		path = "???"
		line = 1
	}
	path = filepath.Base(path)
	fmt.Printf("\t  %s:%d: %s\n", path, line, fmt.Sprint(args...))
	m.t.Fail()
}

func (m Must) Equal(a, b any) {
	if !isEqual(a, b) {
		m.logAndFail(fmt.Sprintf("need %v, got %v", a, b))
	}
}

func (m Must) NotEqual(a, b any) {
	if isEqual(a, b) {
		m.logAndFail(fmt.Sprintf("need %v, got %v", a, b))
	}
}

func (m Must) Nil(a any) {
	m.Equal(nil, a)
}

func (m Must) NotNil(a any) {
	m.NotEqual(nil, a)
}

func (m Must) True(a bool) {
	m.Equal(true, a)
}

func (m Must) False(a bool) {
	m.Equal(false, a)
}

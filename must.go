package must

import (
	"fmt"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"testing"
)

// Must contains single testing.TB instance
// and it will be reuse in same test case
type Must struct {
	t testing.TB
}

// New need testing.TB as parameter
// and Must instance will be returned
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

// Equal check are a and b equal, if not it will be failed
func (m Must) Equal(a, b any) {
	if !isEqual(a, b) {
		m.logAndFail(fmt.Sprintf("need %v, got %v", a, b))
	}
}

// NotEqual reverse of Equal which is failed if a and b are equal
func (m Must) NotEqual(a, b any) {
	if isEqual(a, b) {
		m.logAndFail(fmt.Sprintf("need %v, got %v", a, b))
	}
}

// Nil check is a nil, if not it will be failed
func (m Must) Nil(a any) {
	m.Equal(nil, a)
}

// NotNil reverse of Nil which is failed if a nil
func (m Must) NotNil(a any) {
	m.NotEqual(nil, a)
}

// True check is a true, if not it will be failed
func (m Must) True(a bool) {
	m.Equal(true, a)
}

// False check is a false, if not it will be failed
func (m Must) False(a bool) {
	m.Equal(false, a)
}

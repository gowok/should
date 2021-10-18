package must

import (
	"reflect"
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

func isNil(a interface{}) bool {
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

func isEqual(a, b interface{}) bool {
	if isNil(a) && isNil(b) {
		return true
	}

	if reflect.DeepEqual(a, b) {
		return true
	}

	return reflect.ValueOf(a) == reflect.ValueOf(b)
}

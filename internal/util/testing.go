package util

import (
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Received %v (type %v), expected %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

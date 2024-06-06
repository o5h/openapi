package assert

import (
	"testing"
)

func Nil(t *testing.T, actual any, what ...any) {
	t.Helper()
	if actual != nil {
		logWhat(t, what...)
		t.Fatalf("Expected `nil`, was %v", actual)
	}
}

func NotNil(t *testing.T, actual any, what ...any) {
	t.Helper()
	if actual == nil {
		logWhat(t, what...)
		t.Fatalf("Expected `not nil`, was nil")
	}
}

func Eq[T comparable](t *testing.T, actual T, expected T, what ...any) {
	t.Helper()
	if actual == expected {
		return
	}
	t.Fatalf("Expected =>'%v'<=, was =>'%v'<=", expected, actual)

}

func NotEq[T comparable](t *testing.T, actual T, expected T) {
	t.Helper()
	if actual != expected {
		return
	}
	t.Fatalf("Expected not `%v`, was `%v`", expected, actual)
}

func logWhat(t *testing.T, what ...any) {
	if len(what) > 0 {
		t.Log(what...)
	}
}

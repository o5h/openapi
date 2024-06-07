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
	logWhat(t, what...)
	t.Fatalf("Expected =>'%v'<=, was =>'%v'<=", expected, actual)
}

func NotEq[T comparable](t *testing.T, actual T, expected T, what ...any) {
	t.Helper()
	if actual != expected {
		return
	}
	logWhat(t, what...)
	t.Fatalf("Expected not `%v`, was `%v`", expected, actual)
}

func True(t *testing.T, actual bool, what ...any) {
	t.Helper()
	if actual {
		return
	}
	logWhat(t, what...)
	t.Fatalf("Expected =>'true'<=, was =>'false'<=")
}

func False(t *testing.T, actual bool, what ...any) {
	t.Helper()
	if !actual {
		return
	}
	logWhat(t, what...)
	t.Fatalf("Expected =>'false'<=, was =>'true'<=")
}

func logWhat(t *testing.T, what ...any) {
	if len(what) > 0 {
		t.Log(what...)
	}
}

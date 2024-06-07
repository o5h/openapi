package generator

import (
	"testing"
)

func TestTrimNonLetterPrefix(t *testing.T) {
	type args struct{ s string }
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{s: "12A"}, "A"},
		{"", args{s: "X12A"}, "X12A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimNonLetterPrefix(tt.args.s); got != tt.want {
				t.Errorf("TrimNonLetterPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToPascalCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"AA"}, "AA"},
		{"", args{"A-A"}, "AA"},
		{"", args{"A-Aa"}, "AAa"},
		{"", args{"AAAA-AAAA"}, "AAAAAAAA"},
		{"", args{"Hello World"}, "HelloWorld"},
		{"", args{"HelloWorld"}, "HelloWorld"},
		{"", args{"Hello1 World"}, "Hello1World"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toPascalCase(tt.args.s); got != tt.want {
				t.Errorf("ToPascalCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

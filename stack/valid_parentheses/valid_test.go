package validparentheses

import "testing"

func Test_ValidParentheses(t *testing.T) {
	test := []struct {
		in   string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"[)", false},
		{"({[)", false},
	}
	for _, tt := range test {
		if got := isValid(tt.in); got != tt.want {
			t.Errorf("isValid(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

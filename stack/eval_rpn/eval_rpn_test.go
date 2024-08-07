package evaluatereverse

import (
	"testing"
)

func Test_evalRPN(t *testing.T) {
	cases := []struct {
		tokens []string
		result int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}
	for _, c := range cases {
		res := evalRPN(c.tokens)
		if res != c.result {
			t.Errorf("evalRPN(%v) = %v, but want %v", c.tokens, res, c.result)
		}
	}
}

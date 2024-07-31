package simplifypath

import "testing"

func Test_simplifyPath(t *testing.T) {
	testCases := []struct {
		path     string
		expected string
	}{
		{"/home/", "/home"},
		{"/../", "/"},
	}
	for _, tc := range testCases {
		actual := simplifyPath(tc.path)
		if actual != tc.expected {
			t.Errorf("Given path: %s, Expected: %s, Actual: %s", tc.path, tc.expected, actual)
		}
	}
}

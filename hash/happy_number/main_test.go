package main

import "testing"

func TestStep(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"Test 1", 123, 14},
		{"Test 2", 456, 77},
		{"Test 3", 789, 194},
		{"Test 4", 0, 0},
		{"Test 5", 999, 243},
		{"Test 6", 123456, 91},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := step(tc.input)
			if result != tc.expected {
				t.Errorf("step(%d) = %d; want %d", tc.input, result, tc.expected)
			}
		})
	}
}

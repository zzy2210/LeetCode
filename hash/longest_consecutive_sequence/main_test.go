package main

import (
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got := longestConsecutive(tc.nums)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

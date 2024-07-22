package containsduplicateii_test

import (
	containsduplicateii "leetcode/hash/contains_duplicate_II"
	"testing"
)

func Test_containsNearbyDuplicate(t *testing.T) {
	cases := []struct {
		nums   []int
		k      int
		result bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, false},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, false},
	}
	for _, c := range cases {
		got := containsduplicateii.ContainsNearbyDuplicate(c.nums, c.k)
		if got != c.result {
			t.Errorf("nums: %v, k: %d, got: %v, expect: %v", c.nums, c.k, got, c.result)
		}
	}
}

package summaryranges_test

import (
	summaryranges "leetcode/array/summary_ranges"
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	cases := []struct {
		nums   []int
		result []string
	}{
		{
			nums:   []int{0, 1, 2, 4, 5, 7},
			result: []string{"0->2", "4->5", "7"},
		},
		{
			nums:   []int{0, 2, 3, 4, 6, 8, 9},
			result: []string{"0", "2->4", "6", "8->9"},
		},
		{
			nums:   []int{-1},
			result: []string{"-1"},
		},
		{
			nums:   []int{},
			result: []string{},
		},
	}
	for _, c := range cases {
		got := summaryranges.SummaryRanges(c.nums)
		if !reflect.DeepEqual(got, c.result) {
			t.Errorf("SummaryRanges(%v) = %v; want %v", c.nums, got, c.result)
		}
	}
}

package mergeintervals_test

import (
	mergeintervals "leetcode/array/merge_intervals"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]int
		expect [][]int
	}{
		{"empty", [][]int{}, [][]int{}},
		{"single", [][]int{{1, 3}}, [][]int{{1, 3}}},
		{"merge", [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{"overlap", [][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := mergeintervals.Merge(c.input)
			if !reflect.DeepEqual(got, c.expect) {
				t.Errorf("expected %v, got %v", c.expect, got)
			}
		})
	}
}

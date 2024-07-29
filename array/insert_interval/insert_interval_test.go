package insertinterval

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	cases := []struct {
		name      string
		intervals [][]int
		new       []int
		expected  [][]int
	}{
		{"TestCase1", [][]int{{1, 2}, {3, 5}}, []int{2, 3}, [][]int{{1, 5}}},
		{"TestCase2", [][]int{{1, 2}, {3, 5}}, []int{6, 7}, [][]int{{1, 2}, {3, 5}, []int{6, 7}}},
		{"TestCase6", [][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := insert(c.intervals, c.new)
			if !reflect.DeepEqual(res, c.expected) {
				t.Errorf("Expected: %v, got: %v, case: %v", c.expected, res, c.name)
			}
		})
	}
}

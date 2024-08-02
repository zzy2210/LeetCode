package addtwonumbers

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_add(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	result := &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}}
	test := addTwoNumbers(l1, l2)
	for test != nil && result != nil {
		assert.Equal(t, test.Val, result.Val)
		test = test.Next
		result = result.Next
	}
}

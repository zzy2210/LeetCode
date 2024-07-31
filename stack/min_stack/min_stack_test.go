package minstack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MinStackPush(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	assert.Equal(t, -3, minStack.GetMin())
	assert.Equal(t, -3, minStack.Top())
}

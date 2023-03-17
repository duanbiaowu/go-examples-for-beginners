package creational

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInstance(t *testing.T) {
	pre := NewInstance()
	cur := NewInstance()

	for i := 0; i < 10; i++ {
		assert.Equal(t, pre, cur)
		assert.Same(t, &pre, &cur)
		pre = cur
		cur = NewInstance()
	}
}

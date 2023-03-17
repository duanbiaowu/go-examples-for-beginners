package idiom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Undo(t *testing.T) {
	set := NewIntSet()
	assert.False(t, set.Contains(1))

	set.Add(1)
	assert.True(t, set.Contains(1))

	set.Delete(1)
	assert.False(t, set.Contains(1))
}

func Test_UndoableSet(t *testing.T) {
	set := NewUndoableIntSet()
	assert.False(t, set.Contains(1))

	set.Add(1)
	set.Add(2)
	assert.True(t, set.Contains(1))
	assert.True(t, set.Contains(2))

	err := set.Undo()
	assert.Nil(t, err)
	assert.False(t, set.Contains(2))

	err = set.Undo()
	assert.Nil(t, err)
	assert.False(t, set.Contains(1))
}

func Test_UndoIOC(t *testing.T) {
	set := NewFloatSet()
	assert.False(t, set.Contains(1))

	set.Add(1)
	set.Add(2)
	assert.True(t, set.Contains(1))
	assert.True(t, set.Contains(2))

	err := set.Undo()
	assert.Nil(t, err)
	assert.False(t, set.Contains(2))

	err = set.Undo()
	assert.Nil(t, err)
	assert.False(t, set.Contains(1))
}

package behavioral

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOperation_Operate(t *testing.T) {
	add := Operation{Addition{}}
	assert.Equal(t, add.Operate(10, 20), 30)

	multi := Operation{Multiplication{}}
	assert.Equal(t, multi.Operate(10, 20), 200)
}

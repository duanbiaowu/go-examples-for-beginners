package creational

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGun(t *testing.T) {
	ak47 := NewGun(RifleType)
	assert.Equal(t, ak47, newAK47())

	a92F := NewGun(PistolType)
	assert.Equal(t, a92F, newA92F())
}

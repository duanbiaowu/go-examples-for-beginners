package creational

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBuilder(t *testing.T) {
	assembly := NewBuilder().WithColor(BlueColor)

	familyCar := assembly.WithWheels(SportsWheels).TopSpeed(50 * KPH).Build()
	_ = familyCar.Drive()

	assert.Equal(t, familyCar.Color(), BlueColor)
	assert.Equal(t, familyCar.Wheels(), SportsWheels)
	assert.Equal(t, familyCar.Speed(), 50*KPH)

	sportsCar := assembly.WithWheels(SteelWheels).TopSpeed(150 * KPH).Build()
	_ = sportsCar.Drive()

	assert.Equal(t, sportsCar.Color(), BlueColor)
	assert.Equal(t, sportsCar.Wheels(), SteelWheels)
	assert.Equal(t, sportsCar.Speed(), 150*KPH)
}

package demo

import (
	"testing"
)

func TestInitMission(t *testing.T) {
	mission := InitMission("dj")
	mission.Start()
}

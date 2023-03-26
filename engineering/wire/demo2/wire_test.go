package demo2

import (
	"testing"
)

func TestInitMission(t *testing.T) {
	mission := InitMission("dj", "kitty")
	mission.Start()
}

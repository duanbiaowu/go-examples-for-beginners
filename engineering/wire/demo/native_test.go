package demo

import "testing"

func TestInitMissionNative(t *testing.T) {
	mission := InitMissionNative("dj")
	mission.Start()
}

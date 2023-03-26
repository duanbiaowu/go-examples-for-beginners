package demo3

import (
	"fmt"
	"testing"
)

func TestInitMission(t *testing.T) {
	mission, err := InitMission("dj", "kitty")
	if err != nil {
		fmt.Println(err)
	} else {
		mission.Start()
	}
}

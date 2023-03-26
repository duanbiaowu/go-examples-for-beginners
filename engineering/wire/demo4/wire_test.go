package demo4

import (
	"testing"
)

func TestInitEndingA(t *testing.T) {
	mission := InitEndingA("dj", "kitty")
	mission.Appear()
}

func TestInitEndingB(t *testing.T) {
	mission := InitEndingB("dj", "kitty")
	mission.Appear()
}

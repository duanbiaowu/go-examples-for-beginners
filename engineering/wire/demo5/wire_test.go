package demo5

import (
	"testing"
)

func TestInitEndingA(t *testing.T) {
	mission := InitEndingA("dj")
	mission.Appear()
}

func TestInitEndingB(t *testing.T) {
	mission := InitEndingB("dj")
	mission.Appear()
}

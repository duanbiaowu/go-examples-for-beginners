package creational

import (
	"testing"
)

func TestNewPool(t *testing.T) {
	p := NewPool(10)

	for i := 0; i < 10; i++ {
		<-*p
	}

	close(*p)
}

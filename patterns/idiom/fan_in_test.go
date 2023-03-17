package idiom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FanIn(t *testing.T) {
	cs := make(chan int, 10)
	for i := 0; i < 10; i++ {
		cs <- i
	}

	mergedCh := Merge(cs)
	for i := 0; i < 10; i++ {
		assert.Equal(t, i, <-mergedCh)
	}
	close(cs)
}

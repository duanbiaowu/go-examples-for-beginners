package idiom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FanOut(t *testing.T) {
	ch := make(chan int)
	cs := Split(ch, 10)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for i, c := range cs {
		assert.Equal(t, i, <-c)
	}
}

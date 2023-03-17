package idiom

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CircuitBreaker(t *testing.T) {
	tickets := 2
	fn := func(ctx context.Context) error {
		if tickets == 0 {
			return errors.New("NoTicketsAvailable")
		}
		tickets--
		return nil
	}

	ctx := context.Background()
	circuit := Breaker(fn, 1)

	err := circuit(ctx)
	assert.Nil(t, err)

	err = circuit(ctx)
	assert.Nil(t, err)

	err = circuit(ctx) // ConsecutiveFailures = 1
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("NoTicketsAvailable"), err)

	err = circuit(ctx) // ConsecutiveFailures = 1 && can not retry
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("ErrServiceUnavailable"), err)
}

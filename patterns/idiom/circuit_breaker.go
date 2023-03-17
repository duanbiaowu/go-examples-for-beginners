package idiom

import (
	"context"
	"errors"
	"time"
)

// In the real world, reference: https://github.com/sony/gobreaker

type State int

type Counter interface {
	Count(State)
	ConsecutiveFailures() uint32
	LastActivity() time.Time
	Reset()
}

type Circuit func(ctx context.Context) error

type NormalCounter struct {
	cntMap map[State]uint32
	latAct time.Time
}

const (
	UnknownState State = iota
	FailureState
	SuccessState
)

func (c *NormalCounter) Count(s State) {
	if _, ok := c.cntMap[s]; !ok {
		c.cntMap[s] = 0
	}
	c.cntMap[s]++
}

func (c *NormalCounter) ConsecutiveFailures() uint32 {
	return c.cntMap[FailureState]
}

func (c *NormalCounter) LastActivity() time.Time {
	return c.latAct
}

func (c *NormalCounter) Reset() {
	c.cntMap = make(map[State]uint32)
}

func NewCounter() *NormalCounter {
	return &NormalCounter{
		cntMap: make(map[State]uint32),
		latAct: time.Now(),
	}
}

func Breaker(c Circuit, failureThreshold uint32) Circuit {
	cnt := NewCounter()

	return func(ctx context.Context) error {
		if cnt.ConsecutiveFailures() >= failureThreshold {
			canRetry := func(cnt Counter) bool {
				backoffLevel := cnt.ConsecutiveFailures() - failureThreshold
				shouldRetryAt := cnt.LastActivity().Add(time.Second * 2 << backoffLevel)
				return time.Now().After(shouldRetryAt)
			}

			if !canRetry(cnt) {
				return errors.New("ErrServiceUnavailable")
			}
		}

		if err := c(ctx); err != nil {
			cnt.Count(FailureState)
			return err
		}

		cnt.Count(SuccessState)
		return nil
	}
}

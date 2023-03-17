package idiom

import (
	"errors"
	"time"
)

var (
	ErrNoTickets      = errors.New("semaphore: could not acquire semaphore")
	ErrIllegalRelease = errors.New("semaphore: can't release the semaphore without acquiring it first")
)

type Resource interface {
	Acquire() error
	Release() error
}

type implementation struct {
	sem     chan struct{}
	timeout time.Duration
}

func (s *implementation) Acquire() error {
	select {
	case s.sem <- struct{}{}:
		return nil
	case <-time.After(s.timeout):
		return ErrNoTickets
	}
}

func (s *implementation) Release() error {
	select {
	case <-s.sem:
		return nil
	case <-time.After(s.timeout):
		return ErrIllegalRelease
	}
}

func NewSem(tickets int, timeout time.Duration) Resource {
	return &implementation{
		sem:     make(chan struct{}, tickets),
		timeout: timeout,
	}
}

package tutorial

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvent_DIStart(t *testing.T) {
	message := NewMessage()
	greeter := NewGreeter(message)
	event := NewEvent(greeter)

	event.Start()
}

func TestEvent_WireStart(t *testing.T) {
	event := InitializeEvent()
	event.Start()
}

func TestEvent_WireChangesStart(t *testing.T) {
	event2, err := InitializeEvent2("hello world")
	if err != nil {
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("could not create event: event greeter is grumpy"), err)
	} else {
		event2.Start()
	}
}

func TestEvent_WireChangesWithSetStart(t *testing.T) {
	event2, err := InitializeEvent2WithSet("hello kitty")
	if err != nil {
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("could not create event: event greeter is grumpy"), err)
	} else {
		event2.Start()
	}
}

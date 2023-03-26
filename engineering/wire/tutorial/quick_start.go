package tutorial

import (
	"errors"
	"time"
)

// A First Pass of Building the Greeter Program
//  1) a message for a greeter,
//  2) a greeter who conveys that message, and
//  3) an event that starts with the greeter greeting guests.

type Message string

func NewMessage() Message {
	return Message("Hi there!")
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

type Greeter struct {
	Message Message // <- adding a Message field
}

func (g Greeter) Greet() Message {
	return g.Message
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	println(msg)
}

// Making changes with wire

func NewMessage2(phrase string) Message {
	return Message(phrase)
}

type Greeter2 struct {
	Message Message
	Grumpy  bool
}

func NewGreeter2(m Message) Greeter2 {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter2{m, grumpy}
}

func (g Greeter2) Greet() Message {
	if g.Grumpy {
		return "Go away"
	}
	return g.Message
}

func NewEvent2(g Greeter2) (Event2, error) {
	if g.Grumpy {
		return Event2{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event2{Greeter2: g}, nil
}

type Event2 struct {
	Greeter2 Greeter2 // <- adding a Greeter field
}

func (e Event2) Start() {
	msg := e.Greeter2.Greet()
	println(msg)
}

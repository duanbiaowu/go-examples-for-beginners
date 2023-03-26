//go:build wireinject
// +build wireinject

package tutorial

import "github.com/google/wire"

func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}

func InitializeEvent2(phrase string) (Event2, error) {
	wire.Build(NewEvent2, NewGreeter2, NewMessage2)
	return Event2{}, nil
}

var Event2Set = wire.NewSet(NewGreeter2, NewMessage2)

func InitializeEvent2WithSet(phrase string) (Event2, error) {
	wire.Build(NewEvent2, Event2Set)
	return Event2{}, nil
}

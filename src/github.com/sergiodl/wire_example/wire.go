//+build wireinject

package main

import (
	"github.com/google/go-cloud/wire"
)

func InitializeEvent(phrase string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}

package event

import (
	"context"
	"fmt"
)

type HandleFunc func(context.Context, Event) error

type Event struct {
	Source    string
	Type      string
	RawParams []byte
}

func (e Event) String() string {
	return fmt.Sprintf("Event{Source:%#v}, Type:%#v, RawParams:%#v}", e.Source, e.Type, string(e.RawParams))
}

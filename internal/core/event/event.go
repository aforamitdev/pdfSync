package event

import (
	"context"

	logger "github.com/aforamitdev/pdfsync/zero"
)

type Core struct {
	log      *logger.Logger
	handlers map[string]map[string][]HandleFunc
}

func NewCore(log *logger.Logger) *Core {
	return &Core{log: log, handlers: map[string]map[string][]HandleFunc{}}
}

func (c *Core) SendEvent(ctx context.Context, event Event) error {
	// c.log.Infow("sendevent", "trace_id", "make context call ", "status", "started", "source", event.Source, "type", event.Type, "params", event.RawParams)

	// if m, ok := c.handlers[event.Source]; ok {
	// 	if hfs, ok := m[event.Source]; ok {
	// 		for _,hf:=
	// 	}
	// }

	return nil
}

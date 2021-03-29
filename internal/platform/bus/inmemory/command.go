package inmemory

import (
	"context"
	"log"

	"github.com/jorgechavezrnd/go-hexagonal_http_api/kit/command"
)

// CommandBus is an in-memory implementation of the command.Bus.
type CommandBus struct {
	handlers map[command.Type]command.Handler
}

// NewCommandBus initializes a new instance of CommandBus.
func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[command.Type]command.Handler),
	}
}

// Dispatch implements the command.Bus interface.
func (b *CommandBus) Dispatch(ctx context.Context, cmd command.Command) error {
	handler, ok := b.handlers[cmd.Type()]
	if !ok {
		return nil
	}

	// For Async Command Bus implementation:
	go func() {
		err := handler.Handle(ctx, cmd)
		if err != nil {
			log.Printf("Error while handling %s - %s\n", cmd.Type(), err)
		}
	}()

	return nil

	// For Sync Command Bus implementation:
	// return handler.Handle(ctx, cmd)
}

// Register implements the command.Bus interface.
func (b *CommandBus) Register(cmdType command.Type, handler command.Handler) {
	b.handlers[cmdType] = handler
}

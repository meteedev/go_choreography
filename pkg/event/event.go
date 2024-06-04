package event

import (
	"context"
)

type EventConfig struct {
	Name    string
	Handler func(ctx context.Context, payload []byte) error
}

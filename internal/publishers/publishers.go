package publishers

import (
	"context"

	"github.com/umkh/outbox_pattern/internal/publishers/default_publisher"
)

type Publisher interface {
	Push(ctx context.Context) error
}

type Config struct {
	Default default_publisher.Config
}

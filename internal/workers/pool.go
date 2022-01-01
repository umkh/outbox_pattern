package workers

import "github.com/umkh/outbox_pattern/internal/workers/default_worker"

type Config struct {
	Default default_worker.Config
}

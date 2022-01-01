package outbox

import (
	"github.com/umkh/outbox_pattern/internal/errs"

	"github.com/jmoiron/sqlx"
	"github.com/umkh/outbox_pattern/pkg/publisher"
)

type Outbox struct {
	store      IStore
	publishers map[string]*Publisher
}

func New(db *sqlx.DB) *Outbox {
	return &Outbox{
		store: &store{db: db},
	}
}

func (o *Outbox) SetPublisher(exchange string, pub *publisher.Publisher) *Publisher {
	if pb, ok := o.publishers[exchange]; ok {
		return pb
	}

	pb := &Publisher{}

	o.publishers[exchange] = pb

	return pb
}

func (o *Outbox) GetPublisher(exchange string) (*Publisher, error) {
	if pb, ok := o.publishers[exchange]; ok {
		return pb, nil
	}

	return nil, errs.ErrExchangeNotFound
}

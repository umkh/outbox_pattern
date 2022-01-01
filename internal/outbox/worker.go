package outbox

import (
	"context"
	"log"
	"time"
)

func (o *Outbox) clean(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if err := o.store.Clean(); err != nil {
				log.Println(err)
			}
		}

		time.Sleep(5 * time.Minute)
	}
}

func (o *Outbox) Run(ctx context.Context) error {
	c, cancel := context.WithCancel(ctx)
	defer cancel()

	go o.clean(c)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			log.Println("Todo ...")
		}
		time.Sleep(time.Second)
	}
}

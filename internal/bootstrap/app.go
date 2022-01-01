package bootstrap

import (
	"context"
	"fmt"
	"log"

	"github.com/furdarius/rabbitroutine"
	"github.com/umkh/outbox_pattern/internal/workers"
)

type App struct {
	// db      *sqlx.DB
	rb      *rabbitroutine.Connector
	workers wks
}

func New() *App {
	app := new(App)

	// app.initPostgreSQL()
	app.initRabbitMQ()
	app.initWorkers()

	return app
}

func (app *App) Run(ctx context.Context) {

	go func() {
		url := "amqp://guest:guest@localhost:5672"
		if err := app.rb.Dial(ctx, url); err != nil {
			log.Println(err)
		}
	}()
	fmt.Println("RabbitMQ has been launched")

	for _, wks := range app.workers {
		go func(w workers.Worker) {
			if err := w.Run(ctx); err != nil {
				log.Panic(err)
			}
		}(wks)

		log.Printf("Worker start ...")
	}

	<-ctx.Done()
}

func (app *App) ShutDown() {
	//
}

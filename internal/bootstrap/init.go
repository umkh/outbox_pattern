package bootstrap

import (
	"time"

	"github.com/furdarius/rabbitroutine"
	"github.com/umkh/outbox_pattern/internal/workers"
	"github.com/umkh/outbox_pattern/internal/workers/default_worker"
)

type wks []workers.Worker

// func (app *App) initPostgreSQL() {
// 	dsn := ""

// 	db, err := sqlx.Connect("postgres", dsn)
// 	if err != nil {
// 		log.Panic(err)
// 		return
// 	}

// 	app.db = db
// }

func (app *App) initRabbitMQ() {
	conn := rabbitroutine.NewConnector(rabbitroutine.Config{
		Wait: 2 * time.Second,
	})

	app.rb = conn
}

func (app *App) initWorkers() {
	app.workers = wks{
		default_worker.New(default_worker.Config{
			Exchange:  "outbox",
			Binding:   "outbox.update",
			QueueName: "outbox.messages",
		}, app.rb),
	}
}

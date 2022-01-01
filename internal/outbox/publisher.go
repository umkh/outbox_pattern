package outbox

const (
	StatusPending = iota
	StatusPublished
	StatusDelivered
)

type Publisher struct {
}

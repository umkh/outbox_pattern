package outbox

type IOutbox interface {
	IStore
}

type IStore interface {
	Save() error
	Update() error
	Clean() error
}

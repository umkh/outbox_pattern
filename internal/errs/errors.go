package errs

type Error string

func (e Error) Error() string { return string(e) }

const (
	ErrNoService        = Error("Service not found !")
	ErrExchangeNotFound = Error("Exchange not found !")
)

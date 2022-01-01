package outbox

import "github.com/jmoiron/sqlx"

type store struct {
	db *sqlx.DB
}

func (s *store) Save() error {
	return nil
}

func (s *store) Update() error {
	return nil
}

func (s *store) Clean() error {
	return nil
}

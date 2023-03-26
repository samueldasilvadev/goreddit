package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewStore(dataSource string) (*Store, error) {
	db, err := sqlx.Open("postgres", dataSource)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting database: %w", err)
	}
	return &Store{
		ThreadStore:  &ThreadStore{DB: db},
		PostStore:    &PostStore{DB: db},
		CommentStore: &CommentStore{DB: db},
	}, nil
}

type Store struct {
	*ThreadStore
	*PostStore
	*CommentStore
}

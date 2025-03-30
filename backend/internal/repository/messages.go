package repository

import (
	"database/sql"
)

type MessagesRepository struct {
	db *sql.DB
}

func newMessagesRepo(db *sql.DB) *MessagesRepository {
	return &MessagesRepository{
		db: db,
	}
}


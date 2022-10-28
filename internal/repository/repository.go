package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Repository struct {
	Conn *sql.DB
}

func NewRepository(conn *sql.DB) *Repository {
	return &Repository{Conn: conn}
}

func NewConnection(dsn string) (*sql.DB, error) {
	conn, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	return conn, err
}

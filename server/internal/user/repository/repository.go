package repository

import (
	"database/sql"
	"github.com/bouhartsev/amonic_airlines/server/internal/user"
)

type repository struct {
	conn *sql.DB
}

func NewUserRepository(conn *sql.DB) user.Repository {
	return &repository{conn: conn}
}

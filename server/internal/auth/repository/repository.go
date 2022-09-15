package repository

import (
	"database/sql"
	"github.com/bouhartsev/amonic_airlines/server/internal/auth"
)

type repository struct {
	conn *sql.DB
}

func NewAuthRepository(conn *sql.DB) auth.Repository {
	return &repository{conn: conn}
}

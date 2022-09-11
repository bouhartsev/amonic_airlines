package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLConnection(connString string) (*sql.DB, error) {
	conn, err := sql.Open("mysql", connString)

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}

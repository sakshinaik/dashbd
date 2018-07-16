package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func MySQLConnect(connectionStr string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connectionStr)
	if err != nil {
		return nil, errors.Wrap(err, "MySQL Connection failed")
	}

	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "MySQL db.Ping func call failed")
	}

	return db, nil
}

package models

import (
	"database/sql"

	"github.com/gocraft/dbr"
)

type Category struct {
	ID       int    `db:"ID"`
	Category string `db:"category"`
}

type Categories []Category

func CategoryList(db *sql.DB) ([]Category, error) {
	rs, err := db.Query("SELECT * FROM c")
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	var categories []Category

	for rs.Next() {
		var category Category
		err := rs.Scan(&category.ID, &category.Category)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)

	}

	return categories, nil
}
func CategoryByID(db *sql.DB, id int) (user User, err error) {

	row, err := db.Query("SELECT * FROM categories WHERE ID = ?", id)
	if err != nil {
		return user, err
	}
	defer row.Close()

	n, err := dbr.Load(row, &user)
	if err != nil {
		return user, err
	}
	if n == 0 {
		err = sql.ErrNoRows
	}

	return user, err
}

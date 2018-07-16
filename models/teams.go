package models

import (
	"database/sql"

	"github.com/gocraft/dbr"
)

type Team struct {
	ID   int    `db:"ID"`
	Team string `db:"team"`
}
type Teams []Team

func TeamList(db *sql.DB) ([]Team, error) {
	rs, err := db.Query("SELECT ID, team FROM teams")
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	var teams []Team
	for rs.Next() {
		var t Team
		if err := rs.Scan(&t.ID, &t.Team); err != nil {
			return nil, err
		}
		teams = append(teams, t)

	}

	return teams, nil
}

func TeamByID(db *sql.DB, id int) (team Team, err error) {

	row, err := db.Query("SELECT * FROM teams WHERE ID = ?", id)
	if err != nil {
		return team, err
	}
	defer row.Close()

	n, err := dbr.Load(row, &team)
	if err != nil {
		return team, err
	}
	if n == 0 {
		err = sql.ErrNoRows
	}

	return team, err
}

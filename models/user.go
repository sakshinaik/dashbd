package models

import (
	"database/sql"

	"github.com/gocraft/dbr"
)

type User struct {
	ID        int    `db:"ID"`
	Username  string `db:"username"`
	Firstname string `db:"firstname"`
	Lastname  string `db:"lastname"`
	Email     string `db:"email"`
	Active    int    `db:"active"`
	TeamID    int    `db:"teamID"`
	Team      Team
}

type Users []User

func UserList(db *sql.DB) ([]User, error) {
	rs, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	var users []User

	for rs.Next() {
		var user User
		err := rs.Scan(&user.ID, &user.Username, &user.Firstname, &user.Lastname, &user.Email, &user.Active, &user.TeamID)

		if err != nil {
			return nil, err
		}

		users = append(users, user)

	}

	return users, nil
}
func UserByID(db *sql.DB, id int) (user User, err error) {

	row, err := db.Query("SELECT * FROM users WHERE ID = ?", id)
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
		return user, err
	}
	team, err := TeamByID(db, user.TeamID)
	if err != nil {
		return user, err
	}
	user.Team = team
	return user, err
}

func Login(db *sql.DB, username string) (user User, err error) {

	row, err := db.Query("SELECT ID FROM users WHERE username = ?", username)
	if err != nil {
		return user, err
	}
	defer row.Close()
	var id int
	row.Next()
	err = row.Scan(&id)
	if err != nil {
		return user, err
	}
	user, err = UserByID(db, id)
	return user, err
}

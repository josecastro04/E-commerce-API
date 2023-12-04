package repositories

import (
	"api/src/models"
	"database/sql"
)

type User struct {
	db *sql.DB
}

func NewRepositoryUser(db *sql.DB) *User {
	return &User{db: db}
}

func (u *User) InsertNewUser(user models.User) error {
	statement, err := u.db.Prepare("insert into user (name, email, password) values(?, ?, ?)")
	if err != nil {
		return err
	}

	if _, err = statement.Exec(&user.Name, &user.Email, &user.Password); err != nil {
		return err
	}

	return nil
}

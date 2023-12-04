package repositories

import (
	"api/src/models"
	"database/sql"
	"errors"
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
	defer statement.Close()

	if _, err = statement.Exec(&user.Name, &user.Email, &user.Password); err != nil {
		return err
	}

	return nil
}

func (u *User) SearchUserByID(userID uint64) (models.User, error) {
	row, err := u.db.Query("select * from user where id = ?", userID)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	if row.Next() {
		var user models.User

		if err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedIn); err != nil {
			return models.User{}, err
		}
		return user, nil
	}

	return models.User{}, errors.New("no user info")
}

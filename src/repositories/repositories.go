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

func (u *User) SearchUserByEmail(email string) (models.User, error) {
	row, err := u.db.Query("select * from user where email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	var user models.User
	if row.Next() {
		if err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleType, &user.CreatedIn); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (u *User) InsertNewUser(user models.User) error {
	statement, err := u.db.Prepare("insert into user (name, email, password, roletype) values(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(&user.Name, &user.Email, &user.Password, &user.RoleType); err != nil {
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

	var user models.User
	if row.Next() {
		if err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleType, &user.CreatedIn); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (u *User) SearchPasswordFromUser(userID uint64) (string, error) {
	row, err := u.db.Query("select password from user where id = ?", userID)
	if err != nil {
		return "", err
	}
	defer row.Close()

	var password *string
	if row.Next() {
		if err = row.Scan(&password); err != nil {
			return "", err
		}
	}
	return *password, nil
}

func (u *User) UpdateUserInfo(userID uint64, password models.Password) error {
	statement, err := u.db.Prepare("update user set password = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(&password.New, userID); err != nil {
		return err
	}
	return nil
}

func (u *User) Delete(userID uint64) error {
	statement, err := u.db.Prepare("delete from user where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID); err != nil {
		return err
	}
	return nil
}

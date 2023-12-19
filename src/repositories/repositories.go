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
		if err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Name, &user.Phone, &user.RoleType, &user.CreatedIn); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (u *User) InsertNewUser(user models.User) error {
	statement, err := u.db.Prepare("insert into user (id, username, email, password, name, phone, roletype) values(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(&user.ID, &user.Username, &user.Email, &user.Password, &user.Name, &user.Phone, &user.RoleType); err != nil {
		return err
	}

	return nil
}

func (u *User) InsertDefaultAddress(userID string) error {
	statement, err := u.db.Prepare("insert into address (user_id) value(?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(&userID); err != nil {
		return err
	}

	return nil
}

func (u *User) SearchUserByID(userID string) (models.User, error) {
	row, err := u.db.Query("select u.*, a.city, a.country, a.address, a.postal_code, a.state from user u inner join address a on a.user_id = u.id where u.id = ?", userID)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User
	var city, country, address, postal_code, state sql.NullString
	if row.Next() {
		if err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Name, &user.Phone, &user.RoleType, &user.CreatedIn, &city, &country,
			&address, &postal_code, &state); err != nil {
			return models.User{}, err
		}

		user.Address.City = city.String
		user.Address.Country = country.String
		user.Address.Address = address.String
		user.Address.Postal_Code = postal_code.String
		user.Address.State = state.String
	}
	return user, nil
}

func (u *User) SearchPasswordFromUser(userID string) (string, error) {
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

func (u *User) UpdateUserPassword(userID string, password models.Password) error {
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

func (u *User) Delete(userID string) error {
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

func (u *User) UpdateUserAddress(address models.Address) error {
	statement, err := u.db.Prepare("update address set city = ?, country = ?, address = ?, postal_code = ?, state = ? where user_id = ?")
	if err != nil {
		return err
	}

	if _, err = statement.Exec(&address.City, &address.Country, &address.Address, &address.Postal_Code, &address.State, &address.UserID); err != nil {
		return err
	}

	return nil
}

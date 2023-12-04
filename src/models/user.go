package models

import (
	"api/src/security"
	"errors"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"ID"`
	Name      string    `json:"Name"`
	Email     string    `json:"Email"`
	Password  string    `json:"Password"`
	CreatedIn time.Time `json:"CreatedIn"`
}

func (u *User) Prepare(stage string) error {
	if err := u.valid(); err != nil {
		return err
	}

	if err := u.format(stage); err != nil {
		return err
	}
	return nil
}

func (u *User) valid() error {
	if u.Name == "" {
		return errors.New("the name field cannot be blank and is required")
	}

	if u.Email == "" {
		return errors.New("the email field cannot be blank and is required")
	} else if err := checkmail.ValidateHost(u.Email); err != nil {
		return errors.New("the email needs to be valid")
	}

	if u.Password == "" {
		return errors.New("the password field cannot be blank and is required")
	}

	return nil
}

func (u *User) format(stage string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Password = strings.TrimSpace(u.Password)

	if stage == "signup" {
		hash, err := security.GenerateHashFromPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = string(hash)
	}
	return nil
}

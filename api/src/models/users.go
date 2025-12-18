package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represents a user using the social media
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedIn time.Time `json:"createdIn,omitempty"`
}

// Prepare uses validate and format methods in the received user
func (user *User) Prepare(step string) error {
	if error := user.validate(step); error != nil {
		return error
	}

	if error := user.format(step); error != nil {
		return error
	}
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("name is required and cannot be blank")
	}

	if user.Nick == "" {
		return errors.New("nick is required and cannot be blank")
	}

	if user.Email == "" {
		return errors.New("e-mail is required and cannot be blank")
	}

	if error := checkmail.ValidateFormat(user.Email); error != nil {
		return errors.New("e-mail is invalid")
	}

	if step == "register" && user.Password == "" {
		return errors.New("password is required and cannot be blank")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		hashPassword, error := security.Hash(user.Password)
		if error != nil {
			return error
		}
		user.Password = string(hashPassword)
	}

	return nil
}

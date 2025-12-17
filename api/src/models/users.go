package models

import (
	"errors"
	"strings"
	"time"
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

	user.format()
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

	if step == "register" && user.Password == "" {
		return errors.New("password is required and cannot be blank")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Password = strings.TrimSpace(user.Password)
}

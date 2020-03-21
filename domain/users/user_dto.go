package users

import (
	"strings"

	"gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/errors"
)

const (
	//StatusActive for user
	StatusActive = "active"
)

//User structure
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

//Users marker
type Users []User

//Trim data to make sure not empty values are persisted
func (user *User) Trim() {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)
	user.Status = strings.TrimSpace(user.Status)
}

//Validate the user domain by itself
func (user *User) Validate() errors.RestErr {
	user.Trim()
	user.Email = strings.ToLower(user.Email)
	if user.Email == "" {
		return errors.BadRequestError("invalid email address", errors.New("Invalid e-mail"))
	}
	if user.Password == "" {
		return errors.BadRequestError("invalid password", errors.New("Invalid password"))
	}
	return nil
}

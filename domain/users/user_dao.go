package users

import (
	"fmt"
	"github.com/gasparyanarman/bookstore_users-api/datasources/mysql/users_db"
	"github.com/gasparyanarman/bookstore_users-api/utils/date"
	"github.com/gasparyanarman/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[user.Id]

	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user with id %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]

	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user with id %d already exists", user.Id))
	}

	user.DateCreated = date.GetNowString()
	usersDB[user.Id] = user
	return nil
}

package model

import (
	"github.com/laironacosta/ms-go-layout/internal/domain/user/entity"
)

// User is representing a database model object
type user struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

// NewUser Build the user database model based on the entity model User
func NewUser(userRequest entity.User) *user {
	user := user{
		Name:  userRequest.Name,
		Email: userRequest.Email,
	}

	return &user
}

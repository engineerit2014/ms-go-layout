package models

import (
	"github.com/laironacosta/ms-go-layout/internal/domain/users/entities"
)

// User is representing a database model object
type user struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

// NewUserModel Build the user database model based on the entity model User
func NewUserModel(userRequest entities.User) *user {
	user := user{
		Name:  userRequest.Name,
		Email: userRequest.Email,
	}

	return &user
}

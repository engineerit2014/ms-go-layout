package models

import "github.com/laironacosta/ms-go-layout/internal/app/domain/users/entities"

type user struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

func NewUserModel(userRequest entities.User) *user {
	user := user{
		Name:  userRequest.Name,
		Email: userRequest.Email,
	}

	return &user
}

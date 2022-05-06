package entities

import (
	"github.com/go-playground/validator"
	"github.com/laironacosta/ms-go-layout/internal/app/domain/users/valueobjects"
)

// User entity
type User struct {
	ID     int
	Name   string
	Email  string
	Status vo.UserStatus
}

type UserRequest struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

// NewUser Build the user entity based on userRequest
func NewUser(request UserRequest) (*User, error) {
	// Validate request
	if err := request.Validate(); err != nil {
		return nil, err
	}

	user := User{
		Name:   request.Name,
		Email:  request.Email,
		Status: vo.ActiveUserStatus,
	}

	return &user, nil
}

func (req *UserRequest) Validate() error {
	if err := validator.New().Struct(req); err != nil {
		return err
	}

	return nil
}

// UserIsActive Example of func within an entity
func (r *User) UserIsActive() bool {
	return r.Status == vo.ActiveUserStatus
}

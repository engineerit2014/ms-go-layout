package entity

import (
	"github.com/go-playground/validator"
	"github.com/laironacosta/ms-go-layout/internal/domain/user/valueobject"
)

// User is an entity representing a domain model object (Not a database model) that has identity and is mutable.
type User struct {
	ID     int
	Name   string
	Email  string
	Status vo.UserStatus
}

// UserRequest request to be validated, etc. before creating the entity User
type UserRequest struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

// NewUser Build the user entity based on userRequest
func NewUser(request UserRequest) (*User, error) {
	// Validate UserRequest
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

// Validate is a function that validates the userRequest
func (req *UserRequest) Validate() error {
	if err := validator.New().Struct(req); err != nil {
		return err
	}

	return nil
}

// UserIsActive is a function of the entity that indicates whether the user is active or not.
func (r *User) UserIsActive() bool {
	return r.Status == vo.ActiveUserStatus
}

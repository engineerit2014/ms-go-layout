package createuser

import (
	"context"
)

// InputPort of CreateUser
type InputPort interface {
	Execute(ctx context.Context, request InputPortRequest) (InputPortResponse, error)
}

// InputPortRequest is the request to run the useCase CreateUser with their corresponding json mapper
type InputPortRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// InputPortResponse is the response after running the useCase CreateUser with their corresponding json mapper
type InputPortResponse struct {
	ID int `json:"id"`
}

package createuser

import (
	"context"
)

type InputPort interface {
	Execute(ctx context.Context, request InputPortRequest) (InputPortResponse, error)
}

type InputPortRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type InputPortResponse struct {
	ID int `json:"id"`
}

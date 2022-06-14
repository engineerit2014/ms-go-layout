package repository

import (
	"context"
	"github.com/laironacosta/ms-go-layout/internal/domain/user/entity"
)

// SaveUser repository interface
type SaveUser interface {
	SaveUser(ctx context.Context, user entity.User) (int, error)
}

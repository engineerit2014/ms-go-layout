package repositories

import (
	"context"
	"github.com/laironacosta/ms-go-layout/internal/domain/users/entities"
)

// SaveUserRepo of SaveUser
type SaveUserRepo interface {
	SaveUser(ctx context.Context, user entities.User) (int, error)
}

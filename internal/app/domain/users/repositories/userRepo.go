package repositories

import (
	"context"
	"github.com/laironacosta/ms-go-layout/internal/app/domain/users/entities"
)

type SaveUserRepo interface {
	SaveUser(ctx context.Context, user *entities.User) error
}

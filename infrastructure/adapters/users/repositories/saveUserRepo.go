package repositories

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/laironacosta/ms-go-layout/infrastructure/adapters/users/repositories/models"
	"github.com/laironacosta/ms-go-layout/internal/app/domain/users/entities"
	"github.com/pkg/errors"
	"strings"
)

type saveUserRepository struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) *saveUserRepository {
	return &saveUserRepository{
		db,
	}
}

func (r *saveUserRepository) SaveUser(ctx context.Context, user *entities.User) error {

	// Transform the entity to a database model object and execute the required operations
	userModel := models.NewUserModel(*user)

	if _, err := r.db.Model(userModel).Context(ctx).Insert(); err != nil {
		if pgErr := err.(pg.Error); pgErr != nil && strings.Contains(pgErr.Error(), "duplicate key") {
			return errors.New("ErrorEmailExistsCode")
		}

		return errors.New("ErrorInsertCode")
	}

	user.ID = userModel.ID

	return nil
}

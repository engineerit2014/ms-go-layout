package repositories

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/laironacosta/ms-go-layout/infrastructure/adapters/repositories/models"
	"github.com/laironacosta/ms-go-layout/internal/domain/users/entities"
	"github.com/laironacosta/ms-go-layout/internal/domain/users/repositories"
	"github.com/pkg/errors"
	"strings"
)

type saveUserRepository struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) repositories.SaveUserRepo {
	return &saveUserRepository{
		db,
	}
}

func (r *saveUserRepository) SaveUser(ctx context.Context, user entities.User) (int, error) {
	// Transform the request entity to a database model object and execute the required operations
	userModel := models.NewUserModel(user)

	if _, err := r.db.Model(userModel).Context(ctx).Insert(); err != nil {
		if pgErr := err.(pg.Error); pgErr != nil && strings.Contains(pgErr.Error(), "duplicate key") {
			return 0, errors.New("ErrorEmailExistsCode")
		}

		return 0, errors.New("ErrorInsertCode")
	}

	return userModel.ID, nil
}

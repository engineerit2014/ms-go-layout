package repository

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/laironacosta/ms-go-layout/infrastructure/adapter/repository/model"
	"github.com/laironacosta/ms-go-layout/internal/domain/user/entity"
	"github.com/laironacosta/ms-go-layout/internal/domain/user/repository"
	"github.com/pkg/errors"
	"strings"
)

type saveUser struct {
	db *pg.DB
}

func NewSaveUser(db *pg.DB) repository.SaveUser {
	return &saveUser{
		db,
	}
}

func (r *saveUser) SaveUser(ctx context.Context, user entity.User) (int, error) {
	// Transform the request entity to a database model object and execute the required operations
	userModel := model.NewUser(user)

	if _, err := r.db.Model(userModel).Context(ctx).Insert(); err != nil {
		if pgErr := err.(pg.Error); pgErr != nil && strings.Contains(pgErr.Error(), "duplicate key") {
			return 0, errors.New("ErrorEmailExistsCode")
		}

		return 0, errors.New("ErrorInsertCode")
	}

	return userModel.ID, nil
}

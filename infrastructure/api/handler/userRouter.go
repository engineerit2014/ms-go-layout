package handler

import (
	"github.com/labstack/echo/v4"
)

type UserGroup interface {
	RegisterRouter(*echo.Group)
}

type userGroup struct {
	createUserHandler CreateUserHandler
}

func NewUserGroup(createUserHandler CreateUserHandler) UserGroup {
	return &userGroup{
		createUserHandler,
	}
}

func (r *userGroup) RegisterRouter(group *echo.Group) {
	path := group.Group("/users")
	{
		path.POST("", r.createUserHandler.CreateUserHandler)
	}
}

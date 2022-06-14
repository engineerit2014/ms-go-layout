package router

import (
	"github.com/labstack/echo/v4"
	"github.com/laironacosta/ms-go-layout/infrastructure/router/handler"
)

type UserGroup interface {
	RegisterRouter(*echo.Group)
}

type userGroup struct {
	createUserHandler handler.CreateUser
}

func NewUserGroup(createUserHandler handler.CreateUser) UserGroup {
	return &userGroup{
		createUserHandler,
	}
}

func (r *userGroup) RegisterRouter(group *echo.Group) {
	path := group.Group("/user")
	{
		path.POST("", r.createUserHandler.CreateUserHandler)
	}
}

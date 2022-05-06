package api

import (
	"github.com/labstack/echo/v4"
	"github.com/laironacosta/ms-go-layout/infrastructure/api/users/handler"
)

type router struct {
	server    *echo.Echo
	userGroup handler.UserGroup
}

func NewRouter(
	server *echo.Echo,
	userGroup handler.UserGroup,
) *router {
	return &router{
		server,
		userGroup,
	}
}

func (r *router) Init() {
	basePath := r.server.Group("/ms-go-layout")
	basePath.GET("/health", Health)

	// Register routers
	r.userGroup.RegisterRouter(basePath)

}

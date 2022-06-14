package router

import (
	"github.com/labstack/echo/v4"
	"github.com/laironacosta/ms-go-layout/infrastructure/router/handler"
)

type router struct {
	server    *echo.Echo
	userGroup UserGroup
}

// NewRouter Create a new router
func NewRouter(
	server *echo.Echo,
	userGroup UserGroup,
) *router {
	return &router{
		server,
		userGroup,
	}
}

func (r *router) Init() {
	basePath := r.server.Group("/ms-go-layout")
	basePath.GET("/health", handler.Health)

	// Register routers
	r.userGroup.RegisterRouter(basePath)

}

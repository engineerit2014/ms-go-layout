package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/laironacosta/ms-go-layout/internal/usecase/user/createuser"
	"net/http"
)

type CreateUser interface {
	CreateUserHandler(ctx echo.Context) error
}

type createUser struct {
	inputPort createuser.InputPort
}

func NewCreateUserHandler(inputPort createuser.InputPort) CreateUser {
	return &createUser{
		inputPort,
	}
}

func (h *createUser) CreateUserHandler(c echo.Context) error {
	var req createuser.InputPortRequest
	if err := c.Bind(&req); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Infof("Request received: %+v \n", req)

	resp, err := h.inputPort.Execute(c.Request().Context(), req)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Infof("Response received: %+v \n", resp)
	return c.JSON(http.StatusOK, resp)
}

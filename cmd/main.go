package main

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	pgKit "github.com/laironacosta/kit-go/postgresql"
	"github.com/laironacosta/ms-go-layout/infrastructure/adapters/shared/services"
	"github.com/laironacosta/ms-go-layout/infrastructure/adapters/users/repositories"
	"github.com/laironacosta/ms-go-layout/infrastructure/api"
	"github.com/laironacosta/ms-go-layout/infrastructure/api/users/handler"
	"github.com/laironacosta/ms-go-layout/infrastructure/resources/postgres/migrations"
	"github.com/laironacosta/ms-go-layout/internal/app/usercases/users/createuser"
	"github.com/pkg/errors"
)

// cfg is the struct type that contains fields that stores the necessary configuration
// gathered from the environment.
var cfg struct {
	DBUser string `envconfig:"DB_USER" default:"postgres"`
	DBPass string `envconfig:"DB_PASS" default:"postgres"`
	DBName string `envconfig:"DB_NAME" default:"postgres"`
	DBHost string `envconfig:"DB_HOST" default:"localhost"`
	DBPort int    `envconfig:"DB_PORT" default:"5432"`
	Locale string `envconfig:"LOCALE"  default:"es"`
}

func main() {
	echo := echo.New()
	echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, latency_human={latency_human}\n",
	}))
	echo.Use(middleware.Recover())

	if err := envconfig.Process("LIST", &cfg); err != nil {
		err = errors.Wrap(err, "parse environment variables")
		return
	}

	// Databases
	db := pgKit.NewPgDB(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.DBHost, cfg.DBPort),
		User:     cfg.DBUser,
		Password: cfg.DBPass,
		Database: cfg.DBName,
	})
	migrations.Init(db)

	// Repositories
	userRepo := repositories.NewUserRepository(db)

	// Services
	notificationService := services.NewSendEmailService()

	// Use Cases
	createUserUseCase := createuser.NewUseCase(
		userRepo,
		notificationService,
	)

	// Handlers
	createUserHandler := handler.NewCreateUserHandler(createUserUseCase)

	// Routers
	userGroup := handler.NewUserGroup(createUserHandler)
	r := api.NewRouter(echo, userGroup)
	r.Init()

	echo.Start(":8080")
}

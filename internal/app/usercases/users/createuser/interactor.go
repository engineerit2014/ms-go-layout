package createuser

import (
	"context"
	"github.com/labstack/gommon/log"
	sharedEntities "github.com/laironacosta/ms-go-layout/internal/app/domain/shared/notifications/entities"
	"github.com/laironacosta/ms-go-layout/internal/app/domain/shared/notifications/services"
	"github.com/laironacosta/ms-go-layout/internal/app/domain/users/entities"
	"github.com/laironacosta/ms-go-layout/internal/app/domain/users/repositories"
)

type createUserInteractor struct {
	saveUserRepo     repositories.SaveUserRepo
	sendEmailService services.SendEmailService
}

func NewUseCase(
	saveUserRepo repositories.SaveUserRepo,
	sendEmailService services.SendEmailService,
) InputPort {
	return &createUserInteractor{
		saveUserRepo,
		sendEmailService,
	}
}

// Execute the useCase CreateUser
func (cu *createUserInteractor) Execute(ctx context.Context, request InputPortRequest) (InputPortResponse, error) {
	user, err := entities.NewUser(entities.UserRequest{
		Name:  request.Name,
		Email: request.Email,
	})
	if err != nil {
		return InputPortResponse{}, err
	}

	// Save user into the database
	if err := cu.saveUserRepo.SaveUser(ctx, user); err != nil {
		return InputPortResponse{}, err
	}

	// Send email to the user
	cu.sendEmailService.SendEmail(ctx, sharedEntities.SendEmailServiceRequest{
		EmailDestination: user.Email,
		Subject:          "User creation",
		ContentBody:      "Success",
	})

	if user.UserIsActive() {
		log.Info(user.Status)
	}

	return InputPortResponse{
		ID: user.ID,
	}, nil
}

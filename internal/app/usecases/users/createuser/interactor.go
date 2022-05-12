package createuser

import (
	"context"
	sharedEntities "github.com/laironacosta/ms-go-layout/internal/domain/notifications/entities"
	"github.com/laironacosta/ms-go-layout/internal/domain/notifications/services"
	"github.com/laironacosta/ms-go-layout/internal/domain/users/entities"
	"github.com/laironacosta/ms-go-layout/internal/domain/users/repositories"
)

type createUserInteractor struct {
	saveUserRepo     repositories.SaveUserRepo // Dependency defined as an interface at the domain layer
	sendEmailService services.SendEmailService // Dependency defined as an interface at the domain layer
}

// NewUseCase is constructor for create default implementation of useCase CreateUser
func NewUseCase(
	saveUserRepo repositories.SaveUserRepo,
	sendEmailService services.SendEmailService,
) InputPort {
	return &createUserInteractor{
		saveUserRepo,
		sendEmailService,
	}
}

// Execute createUser useCase responsible for orchestrating the flow of data to and from the entities,
// and directing them to use its services to achieve the use case objectives
func (cu *createUserInteractor) Execute(ctx context.Context, request InputPortRequest) (InputPortResponse, error) {
	// 1. Build User entity from request
	user, err := entities.NewUser(entities.UserRequest{
		Name:  request.Name,
		Email: request.Email,
	})
	if err != nil {
		return InputPortResponse{}, err
	}

	// 2. Save user into the database
	id, err := cu.saveUserRepo.SaveUser(ctx, *user)
	if err != nil {
		return InputPortResponse{}, err
	}

	user.ID = id

	// 3. Send email to the user
	cu.sendEmailService.SendEmail(ctx, sharedEntities.SendEmailServiceRequest{
		EmailDestination: user.Email,
		Subject:          "User creation",
		ContentBody:      "Success",
	})

	// 4. Build and return response
	return InputPortResponse{
		ID: user.ID,
	}, nil
}

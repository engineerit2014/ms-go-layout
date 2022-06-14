package createuser

import (
	"context"
	sharedEntities "github.com/laironacosta/ms-go-layout/internal/domain/notification/entity"
	"github.com/laironacosta/ms-go-layout/internal/domain/notification/service"
	"github.com/laironacosta/ms-go-layout/internal/domain/user/entity"
	"github.com/laironacosta/ms-go-layout/internal/domain/user/repository"
)

type createUserInteractor struct {
	saveUserRepo     repository.SaveUser      // Dependency defined as an interface at the domain layer
	sendEmailService service.SendEmailService // Dependency defined as an interface at the domain layer
}

// NewUseCase is constructor for create default implementation of useCase CreateUser
func NewUseCase(
	saveUserRepo repository.SaveUser,
	sendEmailService service.SendEmailService,
) InputPort {
	return &createUserInteractor{
		saveUserRepo,
		sendEmailService,
	}
}

// Execute createUser useCase responsible for orchestrating the flow of data to and from the entity,
// and directing them to use its service to achieve the use case objectives
func (cu *createUserInteractor) Execute(ctx context.Context, request InputPortRequest) (InputPortResponse, error) {
	// 1. Build User entity from request
	user, err := entity.NewUser(entity.UserRequest{
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

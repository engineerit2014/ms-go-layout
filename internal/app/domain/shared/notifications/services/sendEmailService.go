package services

import (
	"context"
	"github.com/laironacosta/ms-go-layout/internal/app/domain/shared/notifications/entities"
)

type SendEmailService interface {
	SendEmail(ctx context.Context, request entities.SendEmailServiceRequest) error
}

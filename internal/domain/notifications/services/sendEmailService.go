package services

import (
	"context"
	"github.com/laironacosta/ms-go-layout/internal/domain/notifications/entities"
)

// SendEmailService of SendEmail
type SendEmailService interface {
	SendEmail(ctx context.Context, request entities.SendEmailServiceRequest) error
}

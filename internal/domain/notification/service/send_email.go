package service

import (
	"context"
	"github.com/laironacosta/ms-go-layout/internal/domain/notification/entity"
)

// SendEmailService of SendEmail
type SendEmailService interface {
	SendEmail(ctx context.Context, request entity.SendEmailServiceRequest) error
}

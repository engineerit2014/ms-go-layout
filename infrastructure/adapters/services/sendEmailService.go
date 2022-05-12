package services

import (
	"context"
	"github.com/laironacosta/ms-go-layout/internal/domain/notifications/entities"
)

type sendEmailService struct {
	// inject external service to send emails
}

func NewSendEmailService() *sendEmailService {
	return &sendEmailService{}
}

func (r *sendEmailService) SendEmail(_ context.Context, _ entities.SendEmailServiceRequest) error {
	// SendEmail
	return nil
}

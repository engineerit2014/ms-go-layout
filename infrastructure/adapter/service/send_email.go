package service

import (
	"context"
	"github.com/laironacosta/ms-go-layout/internal/domain/notification/entity"
)

type sendEmail struct {
	// inject external service to send emails
}

func NewSendEmail() *sendEmail {
	return &sendEmail{}
}

func (r *sendEmail) SendEmail(_ context.Context, _ entity.SendEmailServiceRequest) error {
	// SendEmail
	return nil
}

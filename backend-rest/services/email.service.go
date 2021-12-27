package service

import "github.com/kabi175/alumini-app-backend/domain/models"

type EmailService struct{}

func NewEmailService() models.EmailService {
	return &EmailService{}
}

func (es *EmailService) Send(string, string, string, string) error {
	panic("implement me")
}

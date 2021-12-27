package service

import (
	"os"

	"github.com/kabi175/alumini-app-backend/domain/models"
)

// OTP Service works through email
type MailOTPService struct {
	mailAddr     string
	mailPassword string
	EmailService models.EmailService
}

// Builder method to construct MailOTPService
func NewMailOTPService() models.OTPService {
	return &MailOTPService{
		mailAddr:     os.Getenv("OTP_MAIL_ADDR"),
		mailPassword: os.Getenv("OTP_MAIL_PASSWORD"),
	}
}

// Generate and Send OTP
// To be implemented
func (as *MailOTPService) GenerateOTP(user *models.User) error {
	panic("implement me")
}

// Update User verification status
// To be implemented
func (as *MailOTPService) VerifyOTP(otp string) error {
	panic("implement me")
}

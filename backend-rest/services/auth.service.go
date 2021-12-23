package service

import (
	"github.com/kabi175/alumini-app-backend/domain/apperrors"
	"github.com/kabi175/alumini-app-backend/domain/models"

	"golang.org/x/crypto/bcrypt"
)

// Provides DefaultAuthService
type DefaultAuthService struct {
	ts models.TokenService
	ur models.UserRepo
}

// Builder function to create DefaultAuthService with dependencies
func NewDefaultAuthService(ts models.TokenService, ur models.UserRepo) models.AuthService {
	return &DefaultAuthService{
		ts: ts,
		ur: ur,
	}
}

// Login authenticates user and returns token
// Used by all Users
func (as *DefaultAuthService) LogIn(user *models.User) (string, error) {
	dbUser, err := as.ur.FindByMail(user.MailID)
	if err != nil {
		return "", err
	}
	if dbUser == nil {
		return "", apperrors.NewNotFoundError("Email not found")
	}
	if !dbUser.Active {
		return "", apperrors.NewUnauthorizedError("Account not activated, contact Admin")
	}
	if !dbUser.Verified {
		return "", apperrors.NewUnauthorizedError("Account not verified")
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.PasswordHash), []byte(user.PasswordHash))
	if err != nil {
		return "", apperrors.NewUnauthorizedError("Invalid password")
	}
	token, err := as.ts.GenerateToken(dbUser)
	return token, err
}

// SignUp creates a new user
// Create default user(student) account
func (as *DefaultAuthService) SignUp(user *models.User) error {
	/*  generate password hash with bcrypt
	 	Update user password hash
		Create user
	*/
	byte, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return apperrors.NewInternalServerError(err.Error())
	}
	user.PasswordHash = string(byte)
	_, err = as.ur.Create(user)
	return err
}

// RemoveAccount removes user account
// User by user (self)  and admin
func (as *DefaultAuthService) RemoveAccount(user *models.User) error {
	err := as.ur.DeleteByID(user.ID)
	return err
}

// Update User Active status (account approval)
// used only by admin
func (as *DefaultAuthService) UpdateActive(user *models.User, isActive bool) error {
	err := as.ur.UpdateActive(user.ID, isActive)
	return err
}

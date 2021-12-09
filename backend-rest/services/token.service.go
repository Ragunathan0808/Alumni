package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kabi175/alumini-app-backend/domain/apperrors"
	"github.com/kabi175/alumini-app-backend/domain/models"
)

type JwtTokenService struct {
	accessToken string
}

func NewJwtTokenService() models.TokenService {
	return &JwtTokenService{
		accessToken: os.Getenv("ACCESS_SECRET"),
	}
}

func (jts *JwtTokenService) GenerateToken(user *models.User) (string, error) {
	atClaims := jwt.MapClaims{}

	atClaims["authorized"] = true
	atClaims["userID"] = user.ID
	atClaims["userEmail"] = user.MailID
	atClaims["userType"] = user.Type
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(jts.accessToken))
	if err != nil {
		return "", apperrors.NewInternalServerError(err.Error())
	}
	return token, nil
}

func (jts *JwtTokenService) DecodeToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jts.accessToken), nil
	})
	if err != nil {
		return nil, apperrors.NewUnauthorizedError(err.Error())
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {

		userID, ok := claims["userID"].(float64)
		if !ok {
			return nil, apperrors.NewInternalServerError("cannot cast claims['userID'] to float64")
		}

		email, ok := claims["userEmail"].(string)
		if !ok {
			return nil, apperrors.NewInternalServerError("cannot cast claims['userEmail'] to float64")
		}

		userType, ok := claims["userType"].(string)
		if !ok {
			return nil, apperrors.NewInternalServerError("cannot cast claims['userType'] to float64")
		}

		user := &models.User{
			ID:     uint(userID),
			MailID: email,
			Type:   userType,
		}
		return user, nil
	}
	return nil, apperrors.NewUnauthorizedError("invalid token")
}

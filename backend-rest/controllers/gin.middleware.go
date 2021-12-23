package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kabi175/alumini-app-backend/domain/apperrors"
	"github.com/kabi175/alumini-app-backend/domain/models"
)

type middleware struct {
	tokenService models.TokenService
}

func (m *middleware) userOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("AUTH")
		if token == "" {
			var err error
			token, err = c.Cookie("AUTH")
			if err != nil {
				c.JSON(500, gin.H{"status": "error", "message": "Internal Server Error"})
			}
		}
		user, err := m.tokenService.DecodeToken(token)
		if err != nil {
			code, msg := apperrors.Extract(err)
			c.JSON(code, gin.H{"status": "error", "message": msg})
		}
		if user.Type != "user" {
			err := apperrors.NewUnauthorizedError("request not allowed for user")
			code, msg := apperrors.Extract(err)
			c.JSON(code, gin.H{"status": "error", "message": msg})
		}
		c.Set("user", user)
		c.Next()
	}
}

func (m *middleware) adminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("AUTH")
		if token == "" {
			var err error
			token, err = c.Cookie("AUTH")
			if err != nil {
				c.JSON(500, gin.H{"status": "error", "message": "Internal Server Error"})
			}
		}
		user, err := m.tokenService.DecodeToken(token)
		if err != nil {
			code, msg := apperrors.Extract(err)
			c.JSON(code, gin.H{"status": "error", "message": msg})
		}
		if user.Type != "admin" {
			err := apperrors.NewUnauthorizedError("request not allowed for user")
			code, msg := apperrors.Extract(err)
			c.JSON(code, gin.H{"status": "error", "message": msg})
		}
		c.Set("user", user)
		c.Next()
	}
}

func (m *middleware) userOrAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("AUTH")
		if token == "" {
			var err error
			token, err = c.Cookie("AUTH")
			if err != nil {
				c.JSON(500, gin.H{"status": "error", "message": "Internal Server Error"})
			}
		}
		user, err := m.tokenService.DecodeToken(token)
		if err != nil {
			code, msg := apperrors.Extract(err)
			c.JSON(code, gin.H{"status": "error", "message": msg})
		}
		if user.Type != "admin" && user.Type != "user" {
			err := apperrors.NewUnauthorizedError("request not allowed for user")
			code, msg := apperrors.Extract(err)
			c.JSON(code, gin.H{"status": "error", "message": msg})
		}
		c.Set("user", user)
		c.Next()
	}
}

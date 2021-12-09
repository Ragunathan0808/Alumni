package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kabi175/alumini-app-backend/domain/apperrors"
	"github.com/kabi175/alumini-app-backend/domain/models"
)

type ginAuthController struct {
	as models.AuthService
}

func AddGinAuthController(as models.AuthService, router *gin.Engine) {
	sub := router.Group("/auth")
	auth := &ginAuthController{
		as: as,
	}

	sub.POST("/login", auth.Login)
	sub.POST("/signup", auth.SignUp)
	sub.DELETE("/deactivate", auth.Deactivate)

}

func (auth *ginAuthController) Login(c *gin.Context) {
	user := models.User{}
	c.ShouldBind(&user)
	token, err := auth.as.LogIn(&user)
	if err != nil {
		code, msg := apperrors.Extract(err)
		c.JSON(code, msg)
		return
	}
	c.SetCookie("AUTH", token, 3600, "/", "", false, false)
	c.Header("AUTH", token)
	c.JSON(200, gin.H{"AUTH": token})
}

func (auth *ginAuthController) SignUp(c *gin.Context) {
	user := models.User{}
	c.ShouldBind(&user)
	err := auth.as.SignUp(&user)
	if err != nil {
		code, msg := apperrors.Extract(err)
		c.JSON(code, msg)
		return
	}
	c.JSON(200, gin.H{"message": "User created successfully"})
}

func (auth *ginAuthController) Deactivate(c *gin.Context) {
	userI := c.Value("user")
	user, ok := userI.(*models.User)
	if !ok {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	err := auth.as.RemoveAccount(user)
	if err != nil {
		code, msg := apperrors.Extract(err)
		c.JSON(code, msg)
		return
	}
	c.JSON(200, gin.H{"message": "User deleted successfully"})
}

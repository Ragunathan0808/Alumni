package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kabi175/alumini-app-backend/domain/apperrors"
	"github.com/kabi175/alumini-app-backend/domain/models"
)

type ginUserController struct {
	us models.UserService
}

func AddGinUserController(us models.UserService, mw *AuthMiddleware, router *gin.Engine) {
	userController := &ginUserController{
		us: us,
	}
	sub := router.Group("/user")
	sub.GET("id/:userID", userController.getByID)
	sub.GET("", userController.get)

}

func (uc *ginUserController) getByID(c *gin.Context) {
	userIDStr := c.Param("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		err = apperrors.NewInternalServerError(err.Error())
		code, msg := apperrors.Extract(err)
		c.JSON(code, gin.H{"status": "error", "message": msg})
	}
	user, err := uc.us.GetByID(uint(userID))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error while getting user by ID",
		})
		return
	}
	c.JSON(200, user)
}

func (uc *ginUserController) get(c *gin.Context) {
	userIDStr := c.Query("id")
	email := c.Query("email")
	studID := c.Query("studID")
	batchStr := c.Query("batch")
	desg := c.Query("desg")
	//	country := c.Query("country")

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		userID = 0
	}
	batch, err := strconv.ParseUint(batchStr, 10, 64)
	if err != nil {
		batch = 0
	}
	user := models.User{
		ID:          uint(userID),
		MailID:      email,
		StudID:      studID,
		Batch:       uint(batch),
		Designation: desg,
	}
	users, err := uc.us.GetUser(&user)
	if err != nil {
		code, msg := apperrors.Extract(err)
		c.JSON(code, gin.H{"status": "error", "message": msg})
	}
	c.JSON(200, users)
}

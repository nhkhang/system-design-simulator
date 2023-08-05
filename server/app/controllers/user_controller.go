package controllers

import (
	"system-design-simulator/app/services/user"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAllUserData(c *gin.Context)
}

type UserControllerImpl struct {
	svc user.UserService
}

func (c *UserControllerImpl) GetAllUserData(c *gin.Context) {
	data := c.svc.GetAllUserData()
	c.JSON(200, data)
}

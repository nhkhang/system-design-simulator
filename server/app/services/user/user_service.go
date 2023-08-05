package user

import "github.com/gin-gonic/gin"

type UserService interface {
	GetAllUser(c *gin.Context)
}

type UserServiceImpl struct {
}

func (s *UserServiceImpl) GetAllUser(c *gin.Context) {
	// TODO: Add logic later
	c.JSON(200, "Hello World")
}

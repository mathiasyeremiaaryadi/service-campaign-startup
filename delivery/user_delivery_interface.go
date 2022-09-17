package delivery

import (
	"github.com/gin-gonic/gin"
)

type UserDelivery interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	GetUserByEmail(c *gin.Context)
	GetUserById(c *gin.Context)
	SaveUserAvatar(c *gin.Context)
}

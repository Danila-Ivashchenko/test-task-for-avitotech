package api

import "github.com/gin-gonic/gin"

type UserHandler interface {
	AddUsers(c *gin.Context)
	DeleteUsers(c *gin.Context)
	GetPersentOfUsersIds(c *gin.Context)
	GetUsersIds(c *gin.Context)
}

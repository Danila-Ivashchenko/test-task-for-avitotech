package api

import "github.com/gin-gonic/gin"

type UserInSegmentHandler interface {
	AddUserToSegments(c *gin.Context)
	AddUsersToSegments(c *gin.Context)
	AddUsersWithLimitOffsetToSegments(c *gin.Context)
	AddPersentOfUsersToSegments(c *gin.Context)
	DeleteUserFromSegments(c *gin.Context)
	GetUserInSegments(c *gin.Context)
	GetUsersInSegment(c *gin.Context)
}

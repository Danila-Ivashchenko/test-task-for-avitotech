package api

import "github.com/gin-gonic/gin"

type HistoryHandler interface {
	GetHistoryOfUser(c *gin.Context)
}

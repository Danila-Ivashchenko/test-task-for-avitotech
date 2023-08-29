package api

import "github.com/gin-gonic/gin"

type SegmentHandler interface {
	AddSegment(c *gin.Context)
	DeleteSegment(c *gin.Context)
	GetAllSegments(c *gin.Context)
	UpdateSegment(c *gin.Context)
}

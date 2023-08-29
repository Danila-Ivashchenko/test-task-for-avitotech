package api

import (
	_ "segment-service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type configer interface {
	GetHttpPort() string
	GetEnv() string
}

type api struct {
	userHandler          UserHandler
	segmentHandler       SegmentHandler
	userInSegmentHandler UserInSegmentHandler
	historyHandler       HistoryHandler
	server               *gin.Engine

	port string
	env  string
}

func New(cfg configer, uh UserHandler, sh SegmentHandler, ush UserInSegmentHandler, h HistoryHandler) *api {
	api := &api{
		port:                 cfg.GetHttpPort(),
		env:                  cfg.GetEnv(),
		userHandler:          uh,
		segmentHandler:       sh,
		userInSegmentHandler: ush,
		historyHandler:       h,
		server:               gin.Default(),
	}
	api.bind()
	return api
}

func (a api) bind() {
	a.server.POST("user", a.userHandler.AddUsers)
	a.server.POST("user/get", a.userHandler.GetUsersIds)
	a.server.DELETE("user", a.userHandler.DeleteUsers)
	a.server.POST("user/get/persent", a.userHandler.GetPersentOfUsersIds)

	a.server.POST("segment", a.segmentHandler.AddSegment)
	a.server.GET("segment/all", a.segmentHandler.GetAllSegments)
	a.server.PUT("segment", a.segmentHandler.UpdateSegment)
	a.server.DELETE("segment", a.segmentHandler.DeleteSegment)

	a.server.POST("user_in_segment/add/one", a.userInSegmentHandler.AddUserToSegments)
	a.server.POST("user_in_segment/add/many", a.userInSegmentHandler.AddUsersToSegments)
	a.server.POST("user_in_segment/add/params", a.userInSegmentHandler.AddUsersWithLimitOffsetToSegments)
	a.server.POST("user_in_segment/add/persent", a.userInSegmentHandler.AddPersentOfUsersToSegments)
	a.server.DELETE("user_in_segment", a.userInSegmentHandler.DeleteUserFromSegments)
	a.server.POST("user_in_segment/get/user", a.userInSegmentHandler.GetUserInSegments)
	a.server.POST("user_in_segment/get/segment", a.userInSegmentHandler.GetUsersInSegment)

	a.server.POST("history", a.historyHandler.GetHistoryOfUser)

	if a.env != "prod" {
		a.server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	a.server.Static("/history", "./history")
}

func (a api) Run() error {
	return a.server.Run(":" + a.port)
}

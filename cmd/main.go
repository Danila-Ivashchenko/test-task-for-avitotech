package main

import (
	"segment-service/internal/adapters/api"
	"segment-service/internal/adapters/handler"
	"segment-service/internal/adapters/storage"
	"segment-service/internal/core/service"
	"segment-service/pkg/client"
	"segment-service/pkg/config"
)

// @title segment-service
// @version 1.0
// @description API Server for work with segments and users

// @host localhost:8080
// @BasePath /

func main() {
	// _ = config.LoadEnv(".env")
	cfg := config.GetConfig()
	psqlClinet := client.NewPostgresClient(cfg)
	userStorage := storage.NewUserStorage(psqlClinet)
	userService := service.NewUserService(userStorage)
	userHandler := handler.NewUserHandler(userService)

	segmentStorage := storage.NewSegmentStorage(psqlClinet)
	segmentService := service.NewSegmentService(segmentStorage)
	segmenthandler := handler.NewSegmentHandler(segmentService)

	historyStorage := storage.NewHistoryStorage(psqlClinet)
	historyService := service.NewHistoryService(historyStorage, cfg)
	historyHandler := handler.NewHistoryHandler(historyService)

	userInSegmentStorage := storage.NewUserInSegmentStorage(psqlClinet)
	userInSegmentService := service.NewUserInSegmentsService(
		userInSegmentStorage,
		userService,
		segmentService,
		historyService,
	)
	userInSegmentHandler := handler.NewUserInSegmentHandler(userInSegmentService)
	api := api.New(
		cfg,
		userHandler,
		segmenthandler,
		userInSegmentHandler,
		historyHandler,
	)
	err := api.Run()
	if err != nil {
		panic("fail to run the server")
	}
}

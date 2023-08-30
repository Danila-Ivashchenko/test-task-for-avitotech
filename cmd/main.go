package main

import (
	"segment-service/internal/adapters/api"
	"segment-service/internal/adapters/handler"
	history_storage "segment-service/internal/adapters/storage/history"
	segment_storage "segment-service/internal/adapters/storage/segment"
	user_storage "segment-service/internal/adapters/storage/user"
	user_in_segment_storage "segment-service/internal/adapters/storage/user_in_segment"

	history_service "segment-service/internal/core/service/history"
	segment_service "segment-service/internal/core/service/segment"
	user_service "segment-service/internal/core/service/user"
	user_in_segment_service "segment-service/internal/core/service/user_in_segment"

	"segment-service/pkg/client"
	"segment-service/pkg/config"
)

// @title segment-service
// @version 1.0
// @description API Server for work with segments and users

// @host localhost:8080
// @BasePath /

func main() {
	// err := config.LoadEnv(".env")
	// if err != nil {
	// 	panic(err)
	// }
	cfg := config.GetConfig()
	psqlClinet := client.NewPostgresClient(cfg)
	userStorage := user_storage.NewUserStorage(psqlClinet)
	userService := user_service.NewUserService(userStorage)
	userHandler := handler.NewUserHandler(userService)

	segmentStorage := segment_storage.NewSegmentStorage(psqlClinet)
	segmentService := segment_service.NewSegmentService(segmentStorage)
	segmenthandler := handler.NewSegmentHandler(segmentService)

	historyStorage := history_storage.NewHistoryStorage(psqlClinet)
	historyService := history_service.NewHistoryService(historyStorage, cfg)
	historyHandler := handler.NewHistoryHandler(historyService)

	userInSegmentStorage := user_in_segment_storage.NewUserInSegmentStorage(psqlClinet)
	userInSegmentService := user_in_segment_service.NewUserInSegmentsService(
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

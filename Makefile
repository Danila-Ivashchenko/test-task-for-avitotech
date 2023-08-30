run:
	@docker-compose up -d

conn_to_db:
	@docker exec -it segment-storage psql -U supervisor -d segments_db

clean:
	@docker stop segment-api;
	@docker rm segment-api;
	@docker kill segment-storage
	@docker rm segment-storage;
	@docker rmi test-task-for-avitotech-api;
gen:
	mockgen -source=internal/core/ports/storage/user.go -destination=internal/mocks/storage/user_mock.go
	mockgen -source=internal/core/ports/storage/segment.go -destination=internal/mocks/storage/segment_mock.go
	mockgen -source=internal/core/ports/storage/user-in-segment.go -destination=internal/mocks/storage/user_in_segment_mock.go
	mockgen -source=internal/core/ports/storage/history.go -destination=internal/mocks/storage/history_mock.go

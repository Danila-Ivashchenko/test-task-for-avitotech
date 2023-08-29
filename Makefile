run:
	@docker-compose up -d

conn_to_db:
	@docker exec -it segment-storage psql -U supervisor -d segments_db

clear:
	@docker rm segment-api;
	@docker rmi test-task-for-avitotech-api;

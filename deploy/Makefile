
## up: starts all containers in the background without forcing build

up:
	@echo "Starting Docker images..."
	docker-compose build
	docker-compose up -d
	@echo "Docker images started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

run:
	@docker-compose down || true
	@docker-compose build
	@docker-compose up -d

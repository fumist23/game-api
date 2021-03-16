APP_CONTAINER_NAME=game-api_app
DB_CONTAINER_NAME=game-db
APP=app
DB=db

build: ## docker build
	docker-compose build --no-cache

run: ## docker up
	docker-compose up -d

stop: ## docker stop
	docker-compose stop

down: ## docker down
	docker-compose down

app: ## app container sh
	docker exec -it $(APP_CONTAINER_NAME) sh

db: ## db container bash
	docker exec -it $(DB_CONTAINER_NAME) bash

logs: ## docker logs 
	docker-compose logs -f

logs/app: ## app container logs
	docker-compose logs -f $(APP)

logs/db: ## db container logs
	docker-compose logs -f $(DB)

help: ## Display this help screen
	@grep -E '^[a-zA-Z/_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'	
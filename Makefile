APP=game-api_app_1
DB=game-db

build: ## docker build
	docker-compose build --no-cache

run: ## docker up
	docker-compose up -d

stop: ## docker stop
	docker-compose stop

down: ## docker down
	docker-compose down

app: ## app container sh
	docker exec -it $(APP) sh

db: ## db container bash
	docker exec -it $(DB) bash

logs: ## docker logs 
	docker-compose logs -f

logs/app: ## app container logs
	docker-compose logs -f app

logs/db: ## db container logs
	docker-compose logs -f db

help: ## Display this help screen
	@grep -E '^[a-zA-Z/_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'	
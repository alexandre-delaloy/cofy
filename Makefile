help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup-env: ## Copy sample files 
	if [ -f app/.env ] ; then mv app/.env app/.env.old ; fi
	cp app/.env.sample app/.env
	if [ -f api/.env ] ; then mv api/.env api/.env.old ; fi
	cp api/.env.sample api/.env

start: ## Up the docker-compose without cache or orphans
	docker-compose up \
	--detach \
	--build \
	--remove-orphans \
	--force-recreate

restart: ## down, then up the docker-compose
	make stop
	make start

stop: ## Down the docker-compose 
	docker-compose down

logs: ## Display logs of your containers 
	docker-compose logs --follow

lint:
	gofmt -s -w -l .

init: ## Initialize the project with all required setup
	make setup-env
	make start
	make logs

.PHONY: help

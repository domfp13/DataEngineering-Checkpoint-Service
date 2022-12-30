# -*- coding: utf-8 -*-
# Created by Luis Enrique Fuentes Plata

SHELL = /bin/bash

include .env
.DEFAULT_GOAL := help

.PHONY: setup
setup: ## 1.-Create Docker Image
	@ echo "********** Building image **********"
	@ docker image build --rm -t ${IMAGE_NAME} .
	@ echo "********** Cleanup **********"
	@ docker image prune -f

.PHONY: bootstrap
bootstrap: ## 2.-Run code server
	@ echo "Creating and Starting services"
	@ $(MAKE) setup
	@ docker-compose -f docker-compose.yml up -d --build --remove-orphans

.PHONY: redis
redis: ## 3.- Jump inside redis server: redis-cli
	@ docker exec -it app-redis-data-store /bin/bash

.PHONY: runner
runner: ## 4- Create python container tester
	@ docker container run --rm -it \
      --name python-runer --network=checkpoint-service-network \
      python:alpine3.14 /bin/ash

.PHONY: clean
clean: ## 5.- Clean Docker
	@ docker-compose down -v --rmi all
	@ docker image prune --filter="dangling=true"

help:
	@ echo "Please use \`make <target>' where <target> is one of"
	@ perl -nle'print $& if m{^[a-zA-Z_-]+:.*?## .*$$}' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-25s\033[0m %s\n", $$1, $$2}'

# Created By Luis Enrique Fuentes Plata

SHELL = /bin/bash
include .env

.DEFAULT_GOAL := help

.PHONY: build
build: ## Build Docker Image
	@ docker image build --rm -t ${IMAGE_NAME} .

.PHONY: run
run: ## Test locally
	@ docker container run --rm -it \
 	  --name checkpoint ${IMAGE_NAME}
 	#@ docker-compose up -d --build

.PHONY: clean
clean: ## (Local): Clean Docker
	@ docker-compose down -v
	@ docker rm $(docker ps -f status=exited -q)
	@ docker rm $(docker ps -f status=created -q)
	@ docker image prune --filter="dangling=true"

.PHONY: runner
runner: ## Create a python runner
	@ docker container run --rm -it \
      --name python-runer --network checkpoint-service-network \
      python:alpine3.14 /bin/ash

help:
	@ echo "Please use \`make <target>' where <target> is one of"
	@ perl -nle'print $& if m{^[a-zA-Z_-]+:.*?## .*$$}' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-25s\033[0m %s\n", $$1, $$2}'

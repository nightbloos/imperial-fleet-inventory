default: help

help:
	@echo "How to use Imperial Fleet Inventory Makefile"
	@echo "make start-db        - Start the database"
	@echo "make start-services  - Start the services"
	@echo ""

env:
	make -C services/http env
	make -C services/spaceship env

start-db:
	docker-compose up -d mysql

start-services:
	docker-compose up -d spaceship http


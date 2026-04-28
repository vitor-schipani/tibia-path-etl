.PHONY: help test up

help:
	@echo "
		help - show this msg
		up - bring compose containers UP
		migrate - perform DB migrations
	"

up:
	docker compose up -d

migrate:
	migrate -path migrations -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable up


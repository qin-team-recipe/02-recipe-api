
server:
	go run cmd/app/main.go

air:
	# mysql.server start; air -c .air.toml
	air -c .air.toml

run:
	docker run recipe-api-c

up:
	docker-compose up

build:
	docker-compose build

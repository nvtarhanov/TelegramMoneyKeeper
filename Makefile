.PHONY:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot	

createdb:
	winpty docker exec -it postgres12 createdb --username=root --owner=root telegramdb	

dropdb:
	winpty docker exec -it postgres12 dropdb --username=root telegramdb

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:alpine	

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/telegramdb?sslmode=disable" -verbose up

migratedown:

	migrate -path db/migration -database "postgresql://root:password@localhost:5432/telegramdb?sslmode=disable" -verbose down


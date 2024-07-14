# SetUp
go:
	go mod download && go mod verify
	go get -u ./...

#GoLang Operations
go-generate:
	go generate ./...

update-packages:
	clear
	@go get -u ${PACKAGE}
	@go mod tidy

## Containers local deploy
db:
	docker compose -f ./compose-db.yaml up -d

api-build-run:
	docker compose -f ./compose.yaml up --build

api-run:
	docker compose -f ./compose.yaml up

services-stop:
	docker compose down

## DB Operations
sqlc-generate:
	sqlc generate -f ./internal/pgstore/sqlc.yaml 

new-migration:
	tern new --migrations ./internal/pgstore/migrations ${MIGRATION_NAME}

migrations:
	clear
	@make db
	@tern migrate --migrations ./internal/pgstore/migrations --config ./internal/pgstore/migrations/tern.conf
	@make services-stop
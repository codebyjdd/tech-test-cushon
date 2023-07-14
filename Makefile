.PHONY: run
run:
	@cd cmd/service; go run main.go

.PHONY: up
up:
	@cd sim; docker-compose up -d

.PHONY: down
down:
	@cd sim; docker-compose down

.PHONY: client
client:
	@docker exec -it sim-db-1 mysql -u root -prootpass cushon

.PHONY: sqlc
sqlc:
	@cd internal/db/sqlc; sqlc generate
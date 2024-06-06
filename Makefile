.PHONY: postgres adminer migrate

postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=secret postgres

adminer:
	docker run --rm -ti --network host adminer

server:
	go run .\cmd\server\main.go

migrate:
	migrate -source file://migrations -database postgres://postgres:12345@localhost/postgres?sslmode=disable up

migrate-down:
	migrate -source file://migrations -database postgres://postgres:12345@localhost/postgres?sslmode=disable down

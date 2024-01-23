run:
	go run .

dev:
	air

tidy:
	go mod tidy -v

deps:
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/postgres

dev-deps:
	go install github.com/cosmtrek/air@latest

postgres:
	docker run --name gorm-postgres \
		-e POSTGRES_PASSWORD=postgres \
		-v gorm-postgres:/var/lib/postgresql/data \
		postgres:16

psql:
	docker exec -it gorm-postgres psql



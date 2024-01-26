run:
	go run .

dev:
	air

tidy:
	go mod tidy -v

deps:
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/postgres
	go get -u github.com/jackc/pgtype

dev-deps:
	go install github.com/cosmtrek/air@latest

postgres:
	docker run -d --rm --name gorm-postgres \
		-e POSTGRES_PASSWORD=postgres \
		-v gorm-postgres:/var/lib/postgresql/data \
		-p 5433:5432 \
		postgres:16

psql:
	docker exec -it gorm-postgres su postgres -c psql



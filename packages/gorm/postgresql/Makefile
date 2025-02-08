POSTGRES_PORT=3307
POSTGRES_DATABASE=gorm
POSTGRES_PASSWORD=gormpass
POSTGRES_IMAGENAME=gorm-postgres

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

postgres-start:
	docker run --name ${POSTGRES_IMAGENAME} \
		-e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
		-v gorm-postgres:/var/lib/postgresql/data \
		-p 5433:5432 \
		postgres:17

postgres-stop:
	docker stop ${POSTGRES_IMAGENAME}

psql:
	docker exec -it ${POSTGRES_IMAGENAME} su postgres -c psql



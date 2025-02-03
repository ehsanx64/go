run:
	go run .

dev:
	air

tidy:
	go mod tidy -v

deps:
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/mysql

dev-deps:
	go install github.com/cosmtrek/air@latest

run-mysql:
	docker run --name gorm-mysql \
		-e MYSQL_ROOT_PASSWORD=gorm-mysql \
		-e MYSQL_DATABASE=gorm \
		-e MYSQL_USER=gormuser \
		-e MYSQL_PASSWORD=gormpass \
		-v gorm-mysql:/var/lib/mysql \
		-p 127.0.0.1:3306:3306 mysql:8.4

mysql-shell:
	docker exec -it gorm-mysql mysql -p

mysql-client:
	mysql -h 127.0.0.1 -u gormuser -p

stop-mysql:
	docker stop gorm-mysql
	docker rm gorm-mysql


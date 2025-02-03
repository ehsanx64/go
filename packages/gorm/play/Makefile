MYSQL_PORT=3307
MYSQL_DATABASE=gorm
MYSQL_USERNAME=gormuser
MYSQL_PASSWORD=gormpass
MYSQL_IMAGENAME=gorm-mysql

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

mysql-start:
	docker run --name ${MYSQL_IMAGENAME} \
		-e MYSQL_ROOT_PASSWORD=${MYSQL_IMAGENAME} \
		-e MYSQL_DATABASE=${MYSQL_DATABASE} \
		-e MYSQL_USER=${MYSQL_USERNAME} \
		-e MYSQL_PASSWORD=${MYSQL_PASSWORD} \
		-v gorm-mysql:/var/lib/mysql \
		-p 127.0.0.1:${MYSQL_PORT}:3306 mysql:8.4

mysql-stop:
	docker stop ${MYSQL_IMAGENAME}
	docker rm ${MYSQL_IMAGENAME}

mysql-shell:
	docker exec -it ${MYSQL_IMAGENAME} mysql -p

mysql-client:
	mysql -h 127.0.0.1 -P ${MYSQL_PORT} -u ${MYSQL_USERNAME} -p


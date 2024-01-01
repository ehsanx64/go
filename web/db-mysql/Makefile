run:
	go run .

tidy:
	go mod tidy -v

run-mysql:
	docker run --name db-mysql \
		-e MYSQL_ROOT_PASSWORD=db-mysql \
		-e MYSQL_DATABASE=main \
		-e MYSQL_USER=main \
		-e MYSQL_PASSWORD=main \
		-v db-mysql:/var/lib/mysql \
		-p 127.0.0.1:3306:3306 mysql:5.7

mysql-shell:
	docker exec -it db-mysql mysql -p

mysql-client:
	mysql -h 127.0.0.1 -u main -p

stop-mysql:
	docker stop db-mysql
	docker rm db-mysql


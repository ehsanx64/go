run:
	go run .

tidy:
	go mod tidy -v

run-mysql:
	docker run --name go-mysql \
		-e MYSQL_ROOT_PASSWORD=go-mysql \
		-v go-mysql:/var/lib/mysql \
		-p 127.0.0.1:3306:3306 mysql:5.7

mysql-shell:
	docker exec -it go-mysql mysql -p

mysql-client:
	mysql -h 127.0.0.1 -u root -p

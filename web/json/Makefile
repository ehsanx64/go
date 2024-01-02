JSON = '{"first_name": "John", "last_name": "Doe", "age": 25}'

dev:
	air

run:
	go run .

deps:
	go get github.com/gorilla/sessions

tidy:
	go mod tidy -v

test-encode:
	@curl http://localhost:8080/encode | jq

test-decode:
	@curl -X POST -d ${JSON} http://localhost:8080/decode

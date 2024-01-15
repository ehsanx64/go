dev:
	air

run:
	go run . 

tidy:
	go mod tidy -v

test:
	# Following fails because field_a is required
	curl -s "http://localhost:8080/adam/not-uuid" | jq

	# Following commands should be ok
	curl -s "http://localhost:8080/adam/9543fa9d-830d-4b6e-a338-4db439cbccf4" | jq

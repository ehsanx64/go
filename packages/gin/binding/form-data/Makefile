dev:
	air

run:
	go run . 

tidy:
	go mod tidy -v

test:
	# Following fails because field_a is required
	curl -s "http://localhost:8080/getb?field_b=field_b_value" | jq

	# Following commands should be ok
	curl -s "http://localhost:8080/getb?field_a=field_a_value&field_b=field_b_value" | jq

	curl -s "http://localhost:8080/getc?field_a=field_a_value&field_c=field_c_value" | jq

	curl -s "http://localhost:8080/getd?field_d=field_d_value&field_x=field_x_value" | jq

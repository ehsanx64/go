COOKIE = "sess=MTcwNDE5MTk4NHxEWDhFQVFMX2dBQUJFQUVRQUFBbF80QUFBUVp6ZEhKcGJtY01Ed0FOWVhWMGFHVnVkR2xqWVhSbFpBUmliMjlzQWdJQUFRPT18GohdJBxqG8n2B__uKQiAVMeBzFeEBdejzQMZ4YHn1x8=; Path=/; Expires=Thu, 01 Feb 2024 10:39:44 GMT; Max-Age=2592000"

dev:
	air

run:
	go run .

deps:
	go get github.com/gorilla/sessions

tidy:
	go mod tidy -v

test-secret:
	curl -s http://localhost:8080/secret

test-secret-cookie:
	curl -s --cookie ${COOKIE} http://localhost:8080/secret

test-login:
	curl -s -I http://localhost:8080/login

test-logout:
	curl http://localhost:8080/logout


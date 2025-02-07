run:
	go run .

dev:
	air

tidy:
	go mod tidy -v

deps:
	export GOPROXY=https://goproxy.io,direct
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/sqlite

dev-deps:
	go install github.com/cosmtrek/air@latest


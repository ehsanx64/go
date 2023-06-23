run:
	go run .

clean:
	go clean

tidy:
	go mod tidy -v

build:
	go build .

build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build .

build-windows-i386:
	GOOS=windows GOARCH=386 go build .
	
build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build .

build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build .

build-linux-i386:
	GOOS=linux GOARCH=386 go build .
	
build-bsd-i386:
	GOOS=freebsd GOARCH=386 go build .
	

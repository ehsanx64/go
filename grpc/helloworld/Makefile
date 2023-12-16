run-server:
	go run server/main.go
run-client:
	go run client/main.go
run-client-gopher:
	go run client/main.go --name=Gopher
protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. \
		--go-grpc_opt=paths=source_relative helloworld/helloworld.proto
tidy:
	go mod tidy -v

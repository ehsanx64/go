run:
	go run .

dev:
	air

tidy:
	go mod tidy -v

deps:
	go get github.com/rabbitmq/amqp091-go


dev-deps:
	go install github.com/cosmtrek/air@latest

rabbitmq:
	# latest RabbitMQ 3.12
	docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.12-management




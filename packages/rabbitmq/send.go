package main

import (
	"log"
	"rabbitmq/lib"
)

func main() {
	body := "Hello world!!!"
	lib.InitRabbitmq()
	lib.Publish(body)
	log.Printf(" [x] Send %s\n", body)
}

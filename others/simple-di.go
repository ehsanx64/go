package main

import "fmt"

type MessageService interface {
	Say() string
}

type MessageServiceData struct {
	Short string
	Long  string
}

type helloService struct {
	MessageServiceData
}

func (hs helloService) SayLong() string {
	return hs.Long
}

func (hs helloService) Say() string {
	return hs.Short
}

func NewHelloService(params MessageServiceData) MessageService {
	return helloService{
		MessageServiceData: params,
	}
}

type goodbyeService struct {
	MessageServiceData
}

func (gs goodbyeService) SayLong() string {
	return gs.Long
}

func (gs goodbyeService) Say() string {
	return gs.Short
}

func NewGoodbyeService(params MessageServiceData) MessageService {
	return goodbyeService{
		MessageServiceData: params,
	}
}

func NewMessageService(config string) MessageService {
	if config == "hello" {
		return NewHelloService(MessageServiceData{
			Long:  "Hello",
			Short: "Hi",
		})
	} else {
		return NewGoodbyeService(MessageServiceData{
			Long:  "Goodbye",
			Short: "Bye",
		})
	}
}

func runService(ms MessageService) string {
	fmt.Printf("%+T\n", ms)
	return ms.Say()
}

func main() {
	fmt.Println(runService(NewMessageService("hello")))
	fmt.Println(runService(NewMessageService("goodbye")))
}

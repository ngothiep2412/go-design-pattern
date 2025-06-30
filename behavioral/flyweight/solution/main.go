package main

import "fmt"

type ChatMessage struct {
	Content string
	Sender  *Sender
}
type Sender struct {
	Name   string
	Age    int
	Avatar []byte
}

type ChatMessageFactory struct {
	cachedFactory map[string]*Sender
}

func (f *ChatMessageFactory) getSender(name string) *Sender {
	return f.cachedFactory[name]
}

func main() {
	chatMessageFactory := ChatMessageFactory{cachedFactory: map[string]*Sender{
		"hi": {
			Name:   "Thiep",
			Age:    20,
			Avatar: make([]byte, 1024*300),
		},
		"good morning": {
			Name:   "Test",
			Age:    20,
			Avatar: make([]byte, 1024*300),
		},
	}}

	fmt.Println([]ChatMessage{
		{
			Content: "Hi",
			Sender:  chatMessageFactory.getSender("Thiep"),
		},
	})
}

package main

import "fmt"

type Sender struct {
	Name   string
	Age    int
	Avatar []byte
}

func main() {
	fmt.Println([]Sender{
		{
			Name:   "Steve",
			Age:    20,
			Avatar: make([]byte, 1024*300),
		},
		{
			Name:   "Thiep",
			Age:    20,
			Avatar: make([]byte, 1024*400),
		},
		{
			Name:   "Test",
			Age:    20,
			Avatar: make([]byte, 1024*300),
		},
	})
}

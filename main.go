package main

import (
	"log"
	"net"
)

type service struct {
}

func NewService() *service {
	return &service{}

}

func (s *service) init() {

}

func (s *service) Start() {
	ln, err := net.Listen("tcp", ":3000")

	if err != nil {
		log.Fatalln(err)
	}

	s.acceptLoop(ln)
}

func (s *service) acceptLoop(ln net.Listener) {
	for {
		_, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	s := NewService()
	s.Start()
}

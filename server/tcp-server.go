package main

import (
	"fmt"
	"net"
)

func NewServer(addr string) *Server {
	return &Server{
		listenAddr: addr,
		quitch:     make(chan struct{}),
		msgch:      make(chan Message, 50),
		clients:    make([]*Client, 0),
		userchan:   make(chan string),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	fmt.Printf("Server started on: %s\n", s.listenAddr)

	go s.AcceptConnectionLoop()

	<-s.quitch
	close(s.msgch)

	return nil
}

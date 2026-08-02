package main

import "net"

type Message struct {
	From    string `json:"from"`
	Payload string `json:"payload"`
}

type Client struct {
	conn              net.Conn
	addr              string
	username          string
	connectedToServer bool
}

// THIS IS TEMPORARY SO MAKE SURE TO REMOVE IT
var UserNumber int = 0

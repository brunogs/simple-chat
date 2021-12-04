package main

import (
	"log"
	"net"
)

func main() {
	s := Server{
		conversation: make(chan Message),
		users:        make(map[string]net.Conn),
	}
	go s.start()
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("unable to start: %s", err.Error())
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err.Error())
			continue
		}
		go s.newClient(conn)
	}
}

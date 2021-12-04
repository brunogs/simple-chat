package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

type Server struct {
	conversation chan Message
	users        map[string]net.Conn
}

func (s *Server) start() {
	for {
		select {
		case message := <-s.conversation:
			for k, u := range s.users {
				if strings.Compare(k, message.user) != 0 {
					u.Write([]byte(message.user + "> " + message.content))
				}
			}
		}
	}
}

func (s *Server) newClient(conn net.Conn) {
	log.Printf("new client: %s", conn.RemoteAddr().String())
	s.users[conn.RemoteAddr().String()] = conn
	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}
		s.conversation <- Message{
			user:    conn.RemoteAddr().String(),
			content: msg,
		}
	}
}

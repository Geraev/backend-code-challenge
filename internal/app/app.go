package app

import (
	"fmt"
	. "github.com/geraev/backend-code-challenge/internal/handlers"
	"log"
	"net"
)

func Run() {
	ln, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Accept connection on port")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Calling handleConnection")
		go HandleConnection(conn)
	}
}

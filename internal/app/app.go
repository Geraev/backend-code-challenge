package app

import (
	"fmt"
	"log"
	"net"

	"github.com/geraev/backend-code-challenge/internal/service"
	"github.com/geraev/backend-code-challenge/internal/storage"
)

func Run() {
	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Accept connection on port")

	storage := mapbased.NewStorage()

	serv := service.NewFriends(storage, listener)

	serv.Start()

}

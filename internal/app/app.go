package app

import (
	"github.com/geraev/backend-code-challenge/internal/protocol"
	"github.com/geraev/backend-code-challenge/internal/service"
	"github.com/geraev/backend-code-challenge/internal/storage"
	"log"
)

func Run() {

	listener, err := protocol.NewTCP("", "3331")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("Accept connection on port")

	storage := mapbased.NewStorage()

	serv := service.NewFriends(storage, listener)

	serv.Start()

}

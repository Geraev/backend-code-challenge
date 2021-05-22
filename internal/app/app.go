package app

import (
	"log"

	"github.com/geraev/backend-code-challenge/internal/protocol"
	"github.com/geraev/backend-code-challenge/internal/service"
	"github.com/geraev/backend-code-challenge/internal/storage"
)

func Run() {
	listener, err := protocol.NewTCP("", "3331")
	if err != nil {
		log.Fatal(err)
	}

	storage := mapbased.NewStorage()

	serv := service.NewFriends(storage, listener)

	serv.Start()

}

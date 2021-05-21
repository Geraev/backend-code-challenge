package service

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/geraev/backend-code-challenge/internal/domain"
	"github.com/geraev/backend-code-challenge/internal/repository"
)

type Friends struct {
	storage repository.IStorage
	listener net.Listener
}

func NewFriends(storage repository.IStorage, listener net.Listener) Friends {
	return Friends{
		storage: storage,
		listener: listener,
	}
}

func (s *Friends) Start()  {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go s.handleConnection(conn)
	}
}

func (s *Friends) broadcastOnlineStatus() {

}

func (s *Friends) handleConnection(conn net.Conn) {
	defer func() {
		fmt.Println("Connection closed")
		conn.Close()
	}()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()

		var CurrentUser domain.User
		err := json.Unmarshal([]byte(message), &CurrentUser)

		if err != nil {
			log.Println("incorrect json:", err)
			continue
		}

		s.storage.PutOrUpdateUser(CurrentUser.ID, CurrentUser.Friends, conn)


		fmt.Println("Message Received:", message)
		newMessage := strings.ToUpper(message)
		conn.Write([]byte(newMessage + "\n"))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
	}
}

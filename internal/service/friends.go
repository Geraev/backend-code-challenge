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
	storage  repository.IStorage
	listener net.Listener
}

type UserStatus struct {
	ID     uint64 `json:"user_id"`
	Status bool   `json:"online"`
}

func NewFriends(storage repository.IStorage, listener net.Listener) Friends {
	return Friends{
		storage:  storage,
		listener: listener,
	}
}

func (s *Friends) Start() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go s.handleConnection(conn)
	}
}

func (s *Friends) broadcastOnlineStatus(userId uint64, status bool) {
	followers := s.storage.Followers(userId)
	usr := UserStatus{
		ID:     userId,
		Status: status,
	}
	for _, item := range followers {
		if conn, ok := s.storage.GetUserConn(item); ok {
			json.NewEncoder(conn).Encode(&usr)
		}
	}
}

func (s *Friends) handleConnection(conn net.Conn) {
	defer func() {
		fmt.Println("Connection closed")
		conn.Close()
	}()

	var currentConnUser uint64
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()

		if strings.ToLower(strings.TrimSpace(message)) == "/quit" {
			break
		}

		var CurrentUser domain.User
		err := json.Unmarshal([]byte(message), &CurrentUser)

		if err != nil {
			log.Println("incorrect json:", err)
			continue
		}

		if status := s.storage.GetUserStatus(CurrentUser.ID); status {
			fmt.Fprintf(conn, "User %d already online", CurrentUser.ID)
			continue
		}

		currentConnUser = CurrentUser.ID
		s.storage.AddUser(CurrentUser.ID, CurrentUser.Friends, conn)
		s.broadcastOnlineStatus(currentConnUser, true)
	}

	defer func() {
		s.storage.SetUserOffline(currentConnUser)
		s.broadcastOnlineStatus(currentConnUser, false)
	}()

	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
	}
}

package handlers

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func HandleConnection(conn net.Conn) {
	defer func() {
		fmt.Println("Connection closed")
		conn.Close()
	}()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println("Message Received:", message)
		newMessage := strings.ToUpper(message)
		conn.Write([]byte(newMessage + "\n"))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
	}
}

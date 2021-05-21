package protocol

import (
	"fmt"
	"net"
)

type TCP struct {
	net.Listener
}

func NewTCP(hostname, port string) (TCP, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", hostname, port))

	return TCP{listener}, err
}

package client

import (
	"io"
	"net"

	"github.com/Elvilius/in-memory-store/internal/config"
)

type TCPClient struct {
	conn net.Conn
}

func NewTCPClient(config config.Config) (*TCPClient, error) {
	conn, err := net.Dial("tcp", config.Network.Address)
	if err != nil {
		return nil, err
	}

	return &TCPClient{
		conn: conn,
	}, nil
}

func (c *TCPClient) Send(message []byte) ([]byte, error) {
	if _, err := c.conn.Write(message); err != nil {
		return nil, err
	}

	response := make([]byte, 4<<10)
	count, err := c.conn.Read(response)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return response[:count], nil
}

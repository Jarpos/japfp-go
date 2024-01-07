package communication

import (
	"fmt"
	"net"
)

type Server struct {
	Host       net.Addr
	connection net.Conn

	SizeX int
	SizeY int
}

func CreateServer(b1 byte, b2 byte, b3 byte, b4 byte, port uint16) Server {
	return Server{
		Host: &net.TCPAddr{
			IP:   net.IPv4(b1, b2, b3, b4),
			Port: int(port),
		},
	}
}

func (s *Server) Connect() error {
	c, err := net.Dial(s.Host.Network(), s.Host.String())
	if err != nil {
		return fmt.Errorf(
			"could not establish connection with server %s. %e", s.Host.String(), err)
	}

	s.connection = c

	x, y, err := GetSize(*s)
	if err != nil {
		return fmt.Errorf(
			"could not get canvas size from server. %e", err)
	}

	s.SizeX = x
	s.SizeY = y

	return nil
}

func (s *Server) Disconnect() {
	s.connection.Close()
}

func (s *Server) read() (string, error) {
	str := ""

	for {
		buffer := make([]byte, 64)
		length, err := s.connection.Read(buffer)
		if err != nil {
			return "", err
		}

		str += string(buffer[:length])
		if str[len(str)-1] == '\n' {
			break
		}
	}

	return str[:len(str)-1], nil
}

func (s *Server) write(payload string) {
	s.connection.Write([]byte(payload))
}

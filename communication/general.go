package communication

import "net"

func read(connection net.Conn) (string, error) {
	s := ""

	for {
		buffer := make([]byte, 64)
		length, err := connection.Read(buffer)
		if err != nil {
			return "", err
		}

		s += string(buffer[:length])
		if s[len(s)-1] == '\n' {
			break
		}
	}

	return s[:len(s)-1], nil
}

func write(connection net.Conn, payload string) {
	connection.Write([]byte(payload))
}

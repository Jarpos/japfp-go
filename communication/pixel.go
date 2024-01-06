package communication

import (
	"fmt"
	"net"
)

func ReadPixel(connection net.Conn, x int, y int) error {
	write(connection, fmt.Sprintf("PX %d %d\n", x, y))

	s, err := read(connection)
	if err != nil {
		return err
	}

	println(s)
	return nil
}

func WritePixel(connection net.Conn, x int, y int, color uint32) {
	write(connection, fmt.Sprintf("PX %d %d %06x\n", x, y, color))
}

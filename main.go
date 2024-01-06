package main

import (
	"Jarpos/japfp-go/communication"
	"fmt"
	"net"
	"time"

	_ "image"
	_ "image/jpeg"
	_ "image/png"
)

const (
	SERVER_HOST = "127.0.0.1"
	SERVER_PORT = "1337"
	SERVER_TYPE = "tcp"
)

func main() {
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	x, y, err := communication.GetSize(connection)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connection to %s:%s established\n", SERVER_HOST, SERVER_PORT)
	fmt.Printf("Canvas size %dx%d (%d pixels)\n", x, y, x*y)

	writeScreen(connection, x, y)
}

func writeScreen(connection net.Conn, x int, y int) {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			communication.WritePixel(
				// connection, i, j, makeColor(uint8(i), uint8(j), uint8(time.Now().UnixMilli())))
				// connection, i, j, rgb)
				// connection, i, j, makeColor64(time.Now().UnixNano(), time.Now().UnixNano(), time.Now().UnixNano()))
				connection, i, j, makeColor64(time.Now().UnixMicro(), time.Now().UnixMicro(), time.Now().UnixMicro()))
			// connection, i, j, makeColor64(time.Now().UnixMilli(), time.Now().UnixMilli(), time.Now().UnixMilli()))
		}
	}
}

func makeColor(r uint8, g uint8, b uint8) uint32 {
	return (uint32(r) << 16) | (uint32(g) << 8) | uint32(b)
}

func makeColor64(r int64, g int64, b int64) uint32 {
	return makeColor(uint8(r), uint8(g), uint8(b))
}

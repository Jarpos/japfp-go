package main

import (
	"Jarpos/japfp-go/communication"
	"fmt"
	"net"
	"os"

	"image"
	"image/color"
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

	img, _ := readImage(os.Args[1])
	fn := func(x int, y int) color.Color {
		return img.At(x%img.Bounds().Max.X, y%img.Bounds().Max.Y)
	}

	writeScreen(connection, x, y, fn)
}

func writeScreen(connection net.Conn, x int, y int, f func(int, int) color.Color) {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			communication.WritePixel(connection, i, j, f(i, j))
		}
	}
}

func readImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	image, _, err := image.Decode(f)
	return image, err
}

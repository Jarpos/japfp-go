package main

import (
	"Jarpos/japfp-go/communication"
	"Jarpos/japfp-go/writer"
	"fmt"
	"os"
	"time"

	"image"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	start := time.Now()
	server := communication.CreateServer(127, 0, 0, 1, 1337)

	err := server.Connect()
	if err != nil {
		panic(err)
	}
	defer server.Disconnect()

	fmt.Printf("Connection to %s established\n", server.Host.String())
	fmt.Printf("Canvas size %dx%d (%d pixels)\n", server.SizeX, server.SizeY, server.SizeX*server.SizeY)

	img, err := readImage(os.Args[1])
	if err != nil {
		panic(err)
	}

	writer.ChanneledTiling(server, img)
	// writer.WriteTiling(server, img)

	elapsed := time.Since(start)
	println("took:", elapsed.Milliseconds(), "ms")
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

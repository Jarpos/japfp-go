package main

import (
	"Jarpos/japfp-go/communication"
	"fmt"
	"os"

	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	server := communication.CreateServer(127, 0, 0, 1, 1337)

	err := server.Connect()
	if err != nil {
		panic(err)
	}
	defer server.Disconnect()

	fmt.Printf("Connection to %s established\n", server.Host.String())
	fmt.Printf("Canvas size %dx%d (%d pixels)\n", server.SizeX, server.SizeY, server.SizeX*server.SizeY)

	img, _ := readImage(os.Args[1])
	fn := func(x int, y int) color.Color {
		return img.At(x%img.Bounds().Max.X, y%img.Bounds().Max.Y)
	}

	writeScreen(server, server.SizeX, server.SizeY, fn)
}

func writeScreen(s communication.Server, x int, y int, f func(int, int) color.Color) {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			communication.WritePixel(s, i, j, f(i, j))
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

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
	if len(os.Args) < 3 {
		printInfo()
		return
	}

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

	writer.FUNCTIONS[os.Args[2]].Writer(server, img)

	elapsed := time.Since(start)
	fmt.Printf("Time taken: %dms (roughly)\n", elapsed.Milliseconds())
	logTimings(&server, elapsed)
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

func printInfo() {
	fmt.Printf("%s [PICTURE] [WRITER]\n\n", os.Args[0])
	fmt.Printf("Writers are:\n")
	for k, v := range writer.FUNCTIONS {
		fmt.Printf("%s: %s\n", k, v.Help)
	}
	fmt.Printf("\n")
}

func logTimings(server *communication.Server, elapsed time.Duration) {
	f, _ := os.OpenFile("timings.log", (os.O_APPEND | os.O_CREATE | os.O_WRONLY), 0644)
	defer f.Close()
	fmt.Fprintf(f, "%d\t%d\t%dms\n", server.SizeX, server.SizeY, elapsed.Milliseconds())
}

package main

import (
	"Jarpos/japfp-go/communication"
	"Jarpos/japfp-go/writer"

	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"image"
	_ "image/jpeg"
	_ "image/png"
)

type Settings struct {
	Writer    string
	ImagePath string
	Server    string
	Port      uint
}

func main() {
	start := time.Now()

	settings := parseArgs()

	server := communication.CreateServer(net.ParseIP(settings.Server), uint16(settings.Port))
	err := server.Connect()
	if err != nil {
		panic(err)
	}
	defer server.Disconnect()

	fmt.Printf("Connection to %s established\n", server.Host.String())
	fmt.Printf("Canvas size %dx%d (%d pixels)\n", server.SizeX, server.SizeY, server.SizeX*server.SizeY)

	img, err := readImage(settings.ImagePath)
	if err != nil {
		panic(err)
	}

	writer.FUNCTIONS[settings.Writer].Writer(server, img)

	elapsed := time.Since(start)
	fmt.Printf("Time taken: %dms (roughly)\n", elapsed.Milliseconds())
	logTimings(&server, elapsed, settings.Writer)
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

func parseArgs() Settings {
	settings := Settings{}
	flag.StringVar(&settings.Writer /*****/, "w", "ct", "Writer to choose")
	flag.StringVar(&settings.ImagePath /**/, "i", "pictures/ConveyorBelt.png", "Image to use")
	flag.StringVar(&settings.Server /*****/, "s", "127.0.0.1", "Server to connect to")
	flag.UintVar(&settings.Port /*********/, "p", 1337, "Port to connect to")

	help := flag.Bool("help", false, "Print help information")
	flag.Parse()

	if *help {
		flag.Usage()
		fmt.Printf("\nWriters are:\n")
		for k, v := range writer.FUNCTIONS {
			fmt.Printf("%s: %s\n", k, v.Help)
		}
		fmt.Printf("\n")
		os.Exit(0)
	}

	return settings
}

func logTimings(server *communication.Server, elapsed time.Duration, writerStr string) {
	f, _ := os.OpenFile("timings.log", (os.O_APPEND | os.O_CREATE | os.O_WRONLY), 0644)
	defer f.Close()
	fmt.Fprintf(f, "%d\t%d\t%dms\t%s\n", server.SizeX, server.SizeY, elapsed.Milliseconds(), writerStr)
}

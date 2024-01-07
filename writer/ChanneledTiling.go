package writer

import (
	"Jarpos/japfp-go/communication"
	"fmt"
	"sync"

	"image"
	"image/color"
)

func ChanneledTiling(s communication.Server, img image.Image) {
	f := func(x int, y int) color.Color {
		return img.At((x % img.Bounds().Max.X), (y % img.Bounds().Max.Y))
	}

	tilesX := (s.SizeX / img.Bounds().Max.X) + 1
	tilesY := (s.SizeY / img.Bounds().Max.Y) + 1

	servers := GetServers(tilesX * tilesY)
	fmt.Printf("Writing %d tiles x=%d, y=%d\n", len(servers), tilesX, tilesY)

	var wg sync.WaitGroup
	wg.Add(len(servers))

	for x := 0; x < tilesX; x++ {
		for y := 0; y < tilesY; y++ {
			go func(wg *sync.WaitGroup, sid int, x, y int) {
				defer wg.Done()
				WriteTile(
					servers[sid], f,
					crect(x*img.Bounds().Dx(), y*img.Bounds().Dy()),
					crect(img.Bounds().Dx(), img.Bounds().Dy()),
				)
			}(&wg, y*tilesX+x, x, y)
		}
	}

	wg.Wait()
}

func WriteTile(
	s communication.Server,
	f func(x int, y int) color.Color,
	offset rect,
	bounds rect) {

	for x := 0; x < bounds.X; x++ {
		for y := 0; y < bounds.Y; y++ {
			communication.WritePixel(s, x+offset.X, y+offset.Y, f(x, y))
		}
	}
}

func GetServers(count int) []communication.Server {
	servers := make([]communication.Server, count)

	for i := 0; i < count; i++ {
		servers[i] = communication.CreateServer(127, 0, 0, 1, 1337)
		err := servers[i].Connect()
		if err != nil {
			panic(err)
		}
	}

	return servers
}

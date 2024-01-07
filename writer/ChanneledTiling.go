package writer

import (
	"Jarpos/japfp-go/communication"
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

	var wg sync.WaitGroup
	sid := 0
	servers := GetServers((tilesX * tilesY) + 1)

	for i := 0; i < tilesX; i++ {
		x := i
		for j := 0; j < tilesY; j++ {
			y := j
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				WriteTile(
					servers[id], f,
					crect(x*img.Bounds().Dx(), y*img.Bounds().Dy()),
					crect(img.Bounds().Dx(), img.Bounds().Dy()),
				)
			}(sid)
			sid++
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

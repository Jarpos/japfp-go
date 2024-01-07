package writer

import (
	"Jarpos/japfp-go/communication"

	"image"
	"image/color"
)

func WriteTiling(s communication.Server, i image.Image) {
	f := func(x int, y int) color.Color {
		return i.At(x%i.Bounds().Max.X, y%i.Bounds().Max.Y)
	}

	for i := 0; i < s.SizeX; i++ {
		for j := 0; j < s.SizeY; j++ {
			communication.WritePixel(s, i, j, f(i, j))
		}
	}
}

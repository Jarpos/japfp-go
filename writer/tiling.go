package writer

import (
	"Jarpos/japfp-go/communication"

	"image"
	"image/color"
)

func WriteTiling(s communication.Server, i image.Image) error {
	f := func(x int, y int) color.Color {
		return i.At(x%i.Bounds().Max.X, y%i.Bounds().Max.Y)
	}

	for x := 0; x < s.SizeX; x++ {
		for y := 0; y < s.SizeY; y++ {
			communication.WritePixel(s, x, y, f(x, y))
		}
	}

	return nil
}

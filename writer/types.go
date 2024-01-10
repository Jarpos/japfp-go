package writer

import (
	"Jarpos/japfp-go/communication"

	"image"
)

type Writer interface {
	Write(communication.Server, image.Image)
}

type rect struct {
	X int
	Y int
}

func crect(x int, y int) rect {
	return rect{X: x, Y: y}
}

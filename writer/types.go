package writer

import (
	"Jarpos/japfp-go/communication"

	"image"
)

type WriterFunc = func(communication.Server, image.Image) error

type Writer struct {
	Writer WriterFunc
	Help   string
}

type rect struct {
	X int
	Y int
}

func crect(x int, y int) rect {
	return rect{X: x, Y: y}
}

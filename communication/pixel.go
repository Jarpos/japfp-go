package communication

import (
	"fmt"
	"image/color"
)

func ReadPixel(s Server, x int, y int) error {
	s.write(fmt.Sprintf("PX %d %d\n", x, y))

	str, err := s.read()
	if err != nil {
		return err
	}

	println(str)
	return nil
}

func WritePixel(s Server, x int, y int, color color.Color) {
	s.write(fmt.Sprintf("PX %d %d %06x\n", x, y, makeColor(color)))
}

func makeColor(color color.Color) uint32 {
	r, g, b, _ := color.RGBA()
	return (uint32(uint8(r)) << 16) | (uint32(uint8(g)) << 8) | uint32(uint8(b))
}

// func WritePixel(connection net.Conn, x int, y int, color color.Color) {
// 	write(connection, fmt.Sprintf("PX %d %d %08x\n", x, y, makeColor(color)))
// }
//
// func makeColor(color color.Color) uint32 {
// 	r, g, b, a := color.RGBA()
// 	return (uint32(r) << 24) | (uint32(g) << 16) | (uint32(b) << 8) | (uint32(a))
// }

package communication

import (
	"fmt"
	"strconv"
	"strings"
)

func GetSize(s Server) (int, int, error) {
	s.write("SIZE\n")

	size, err := s.read()
	if err != nil {
		return -1, -1, err
	}

	return parseSize(size)
}

func parseSize(s string) (int, int, error) {
	ss := strings.Split(s, " ")
	if len(ss) != 3 || ss[0] != "SIZE" {
		return -1, -1, fmt.Errorf(
			"expected format: SIZE X Y\\n, got: %s", strconv.QuoteToASCII(s))
	}

	x, e1 := strconv.Atoi(ss[1])
	if e1 != nil {
		return -1, -1, fmt.Errorf(
			"problem converting x value, got: %s", strconv.QuoteToASCII(ss[1]))
	}

	y, e2 := strconv.Atoi(ss[2])
	if e2 != nil {
		return -1, -1, fmt.Errorf(
			"problem converting y value, got: %s", strconv.QuoteToASCII(ss[2]))
	}

	return x, y, nil
}

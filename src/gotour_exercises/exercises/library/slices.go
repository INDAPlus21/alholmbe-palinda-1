package library

import (
	"math/rand"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)

	for i := range s {
		s[i] = make([]uint8, dx)
		for j := range s[i] {
			s[i][j] = uint8(rand.Int())
		}
	}

	return s
}

func RunPic() {
	pic.Show(Pic)
}

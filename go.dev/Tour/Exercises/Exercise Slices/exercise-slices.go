package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		var imageRow = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			imageRow[j] = uint8(i*j)
		}
		image[i] = imageRow
	}
	return image
}

func main() {
	pic.Show(Pic)
}

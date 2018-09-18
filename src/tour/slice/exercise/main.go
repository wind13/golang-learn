package main

import "tour/slice/exercise/pic"

func calc(x, y int) uint8 {
	return uint8((x*x+y*y))
}

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy, dy)
	for i := range img {
		img[i] = make([]uint8, dx, dx)
		for j := range img[i] {
			img[i][j] = calc(i, j)
		}
	}
	return img
}

func main() {
	pic.Show(Pic)
}

package main

// 练习：slice

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, 0, dy)

	for y := 0; y < dy; y++ {
		row := make([]uint8, 0, dx)
		for x := 0; x < dx; x++ {
			row = append(row, uint8(x*x + y*y))
		}
		pic = append(pic, row)
	}

	return pic
}

func main() {
	pic.Show(Pic)
}

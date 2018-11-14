package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	x, y := robotgo.GetMousePos()
	fmt.Println("pos: ", x, y)
	color := robotgo.GetPixelColor(x, y)
	fmt.Println("color----", color)
}

package main

import (
	"fmt"
	"image/color"
)

type Point struct {
	X, Y float64
}

type ColorPoint struct {
	Point
	Color color.RGBA
}

func main() {
	var cp ColorPoint
	// 直接赋值
	cp.X = 1
	fmt.Println(cp.Point.X)
	// 直接给ColorPoint中的Point熟悉赋值
	cp.Point.Y = 2
	fmt.Println(cp.Y)
}

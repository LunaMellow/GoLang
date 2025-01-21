package main

import (
	"GoLang/Exercises/2/geometry"
)

func main() {
	circle := geometry.Circle{Radius: 20}
	geometry.PrintArea(circle)

	square := geometry.Square{Length: 11}
	geometry.PrintArea(square)
}

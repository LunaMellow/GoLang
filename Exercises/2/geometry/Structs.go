package geometry

import (
	"fmt"
	"math"
	"reflect"
)

type Area interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Square struct {
	Length float64
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

func PrintArea(a Area) {
	typeName := reflect.TypeOf(a).Name()
	fmt.Printf("Area of %s: %.2f\n", typeName, a.Area())
}

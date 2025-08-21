package FourtModule

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}
type Rectangle struct {
	Length float64
	Width  float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (c Rectangle) Area() float64 {
	return c.Width * c.Length
}

func first() {
	var c, r Shape = Circle{Radius: 4}, Rectangle{Length: 4, Width: 4}

	fmt.Printf("Circle area: %f", c.Area())
	fmt.Printf("\nRectangle area: %f", r.Area())

}

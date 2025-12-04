package main
import (
	"fmt"
)
/*
	Interface is a collection of method signatures/declarations without definitions.
*/
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64{
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64{
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64{
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64{
	return 2 * 3.14 * c.Radius
}

func PrintArea(s Shape) float64 {
	return s.Area()
}
func PrintPerimeter(s Shape) float64 {
	return s.Perimeter()
}

func main() {
	rectangle := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 5}

	fmt.Println("Area of rectangle: ", PrintArea(rectangle))
	fmt.Println("Area of Circle: ", PrintArea(circle))

	fmt.Println("Perimeter of rectangle: ",PrintPerimeter(rectangle))
	fmt.Println("Perimeter/Circumfrence of Circle: ",PrintPerimeter(circle))
}
package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
	diameter float64
	circumference float64
	area float64
}

func calArea(c Circle) {
	c.area = math.Pi * c.radius * c.radius
	fmt.Println("Area: ",c.area)
}
// passing as reference 
func calArea2(c *Circle) {
	c.area = math.Pi * c.radius * c.radius // Go's pointer dereferencing automatically refers the base one
	fmt.Println("Area by ref: ",(*c).area)
}
func calDiameter(c Circle) {
	c.diameter = 2 * c.radius
	fmt.Println("Diameter: ",c.diameter)
}
func calCircumference(c Circle) {
	c.circumference = 2 * math.Pi * c.radius
	fmt.Println("circrumfrence: ",c.circumference)
}

func main() {
	c := Circle{radius: 5}
	calArea(c)
	calDiameter(c)
	calCircumference(c)
	fmt.Println(c)
	fmt.Printf("%+v\n", c)
	calArea2(&c)
	fmt.Printf("%+v\n", c)
}

package main;

import (
	"fmt"
);

const PI float64 = 3.14159; // constant declaration with explicit type (typed constant)
const appName = "CircleAreaCalculator" // constant declaration with implicit type (untyped constant)

func circleArea(radius float64) float64 {
	fmt.Println("Using", appName);
	area := PI * radius * radius;
	return area;
}

func main() {
	var radius float64;
	fmt.Print("Enter the radius of the circle: ");
	fmt.Scanf("%f", &radius);
	Area := circleArea(radius);

	fmt.Printf("Area of the circle with radius %.2f is %.2f\n", radius, Area);
}

/*
	In Go, constants are declared using the const keyword.
	Constants are immutable, meaning their values cannot be changed after they are defined.

	Constant Declaration Syntax:
	const constantName dataType = value

	Example:
	const pi float64 = 3.14159
	const appName string = "MyApp"

	Constants can be of types: boolean, numeric (int, float64), and string.
*/
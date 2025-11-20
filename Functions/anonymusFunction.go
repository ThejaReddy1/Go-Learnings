package main;

import (
	"fmt"
)

/*
	Anonymous Functions:
	Anonymous function is a function that is declared without any named identifier to refer to it. 
	They can accept inputs and returns outputs, just like standard functions do.
	They can be used for containing functionality that need not to be named and possibly for short-term use.
*/

func main() {

	x := func(l int, b int) int {
		return l * b
	}
	fmt.Printf("%T \n", x) // prints the type of the anonymous function
	fmt.Println("Area of Rectangle is: ", x(10, 5))

	y := func(a int, b int) int {
		return a+b
	}(10,20) // Immediately Invoked Function Expression (IIFE)

	fmt.Printf("%T \n", y) // prints the type of the anonymous function
	fmt.Println("Sum is: ", y)
}
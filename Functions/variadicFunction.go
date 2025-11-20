package main;

import (
	"fmt" 
)
/*
	Variadic Functions:
	Variadic functions are functions that can accept a variable number of arguments.
	They are defined using an ellipsis (...) before the type of the last parameter in the function signature.
	This allows you to pass zero or more arguments of that type when calling the function.
	Inside the function, the variadic parameter is treated as a slice of that type.

	Syntax:
	func functionName(parameter1 type1, parameter2 ...type2) {
	    // Function body
	}
*/
func sum(numbers ...int) (sum int) {
	sum = 0;
	for _, number := range numbers {
		sum += number;
	}
	return
}

func main() {
	total1 := sum(1, 2, 3)
	fmt.Println("Sum of 1, 2, 3 is: ", total1)

	total2 := sum(10, 20, 30, 40, 50)
	fmt.Println("Sum of 10, 20, 30, 40, 50 is: ", total2)
}

// '__' blank identifier is used to ignore the returned value from a function when you don't need it.
/*
	example:
	func divide(x int, y int) (int, int) {
		quotient := x / y
		remainder := x % y
		return quotient, remainder
	}

	func main() {
		q, _ := divide(10, 3) // Ignoring the remainder
		fmt.Println("Quotient: ", q)
	}
*/
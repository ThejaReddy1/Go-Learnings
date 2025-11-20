package main;

import (
	"fmt"
)
// Sample Greet Function
func greet() {
	fmt.Println("Welcome to Go Learnings")
}

func main() {
	greet()
}
/*
	Functions are self-contained blocks of code that perform a specific task.
	They allow you to break down your code into smaller, reusable units.
	Functions can take parameters (input values) and return values (output).
	They enhance code modularity, readability, and reusability.
	Functions are defined using the func keyword, followed by the function name, parameters, and return type (if any).
	Syntax:
	func functionName(parameter1 type1, parameter2 type2) returnType {
	    // Function body
	}

	==>> Function Parameters are the names listed in the function's definition.
	==>> Function Arguments are the actual values passed to the function when it is called.

	return Statement:
	The return statement in Go is used to exit a function and optionally pass back one or more values to the caller.
	It can be used to return multiple values by separating them with commas.
	When a return statement is executed, the function terminates immediately, and control is passed back to the calling function.
	
	There are different types of functions in Go:
	1. Standard Functions: Regular functions that perform specific tasks.
	2. Recursive Functions: Functions that call themselves to solve a problem.
	3. Anonymous Functions: Functions without a name, often used as arguments to other functions.
	4. Variadic Functions: Functions that can accept a variable number of arguments.
	5. Higher-Order Functions: Functions that can take other functions as arguments or return them as results.

	Defer Statement:
	The defer statement in Go is used to ensure that a function call is performed later in a program's execution, usually for purposes of cleanup.
	Deferred functions are executed in LIFO (Last In, First Out) order just before the surrounding function returns.
	Defer is commonly used to close files, release resources, or unlock mutexes.

	
*/
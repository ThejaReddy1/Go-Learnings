package main

import "fmt"

/*
	Passing by value functions:

	-> Function is called by directly passing the value of the variable as an argument.
	-> the parameter is copied into another location of your memory.
	-> So, when accessing or modifying the variable within your function, only the copy is accessed or modified, and the original value is never modified.
	-> All basic types (int, float, bool, string) are passed by value.
	Note that passing by value is usually how we pass values to functions.
*/

func modify( a int) {
	a+=100
	fmt.Println("Address of a inside modify function:", &a) // Output: Address of a inside modify function: 0xc0000140b0
	fmt.Println("Value of a inside modify function:", a)      // Output: Value of a inside modify function: 150
}

func main() {
	a := 50
	modify(a)
	fmt.Println("Address of a in main function:", &a) // Output: Address of a in main function: 0xc0000140a8
	fmt.Println("Value of a after modify function:", a) // Output: Value of a after modify function: 50
}


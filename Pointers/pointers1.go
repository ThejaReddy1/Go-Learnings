package main

import "fmt"

/*
	Pointers in Go

	pointers are variables that store the memory address of another variable.
	They are used to directly manipulate the value at that address.

	They point where the meory is allocated and provide ways to find or even change the value located at the memory location.

	To declare a pointer, use the * operator before the type.
	To get the address of a variable, use the & operator.
	To dereference a pointer (access the value at the address), use the * operator again.

	Example:
	var x int = 42
	var p *int = &x


*/

func main() {
	var x int = 42
	var ptr *int = &x

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value of ptr (address of x):", ptr)
	fmt.Println("Value at the address stored in ptr:", *ptr)
}
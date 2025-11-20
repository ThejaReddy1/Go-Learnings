package main;

import "fmt"

/*
	& operator - The address of a varible can be address obtain by preceding the nale of a variable with an ampersand sign (&), known as address-of operator.

	* operator - It is known as the dereference operator, When placed before an address, it returns the value at that address.

*/

func main() {
	x := 77
	ptr := &x

	fmt.Printf("%T %v \n", &x, &x)
	fmt.Printf("%T %v \n", *(&x), *(&x))
	fmt.Printf("%T %v \n", ptr, *ptr)
}
package main;

import (
	"fmt"
)

func main() {
	var a, b int = 10, 5
	fmt.Println("Initial values: a =", a, ", b =", b)

	// Addition assignment
	a += b
	fmt.Println("a += b:", a)

	// Reset and show subtraction assignment
	a = 10
	a -= b
	fmt.Println("a -= b:", a)

	// Reset and show multiplication assignment
	a = 10
	a *= b
	fmt.Println("a *= b:", a)

	// Reset and show division assignment
	a = 10
	a /= b
	fmt.Println("a /= b:", a)

	// Reset and show modulus assignment
	a = 10
	a %= b
	fmt.Println("a %= b:", a)
}

/*
Assignment Operators in Go:
1. =   : Simple assignment
2. +=  : Addition assignment
3. -=  : Subtraction assignment
4. *=  : Multiplication assignment
5. /=  : Division assignment
6. %=  : Modulus assignment

These operators are used to assign values to variables in Go.
The +=, -=, *=, /=, and %= operators perform the operation on the variable and then assign the result back to that variable.
*/
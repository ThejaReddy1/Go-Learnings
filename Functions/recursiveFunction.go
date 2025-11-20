package main;

import (
	"fmt"
)
/*
	Recursive Functions:
	Recursion is a concept where a function calls itself by direct or indirect means to solve a problem.
	The function keeps calling itself until it reaches a base case, which is a condition that stops the recursion.
	Used to solve a problem where the solution is dependent on the smaller instances of the same problem.
*/
func factorial (n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * factorial(n -1)
	}
}

func main() {
	n := 10
	result := factorial(n)
	fmt.Printf("Factorial of %d is: %d\n", n, result)
}
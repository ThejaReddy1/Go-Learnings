package main;

import (
	"fmt"
)

func main() {
	var a int = 10;
	var b int = 33;

	sum := a + b;  				// addition Operator
	subtraction := a - b;		// subtraction Operator	
	multiplication := a * b;	// multiplication Operator
	division := b / a;		// division Operator
	modulus := b % a;		// modulus Operator
	a++;		// increment Operator
	b--; 		// decrement Operator

	fmt.Println("Sum of ", a, "and", b, "is :", sum)
	fmt.Println("Subtraction of ", a, "and", b, "is :", subtraction)
	fmt.Println("Multiplication of ", a, "and", b, "is :", multiplication)
	fmt.Println("Division of ", b, "by", a, "is :", division)
	fmt.Println("Modulus of ", b, "and", a, "is :", modulus)
	fmt.Println("Incrementing a gives :", a)
	fmt.Println("Decrementing b gives :", b)

	// For String Addition Operator acts as concatenation operator

	var firstName string = "John "
	var lastName string = "Doe"

	fullName := firstName + lastName; // string concatenation

	fmt.Println("Full Name is :", fullName)
}

/*
	increment and decrement operators are unary operators and can only be used in standalone statements.
	`/` Division Operator returns quotient without remainder.
	`%` Modulus Operator returns remainder after division.
*/
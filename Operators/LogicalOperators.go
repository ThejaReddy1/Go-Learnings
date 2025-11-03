package main;

import (
	"fmt"
)

func main() {
	a := true
	b := false

	andResult := a && b
	orResult := a || b
	notA := !a
	notB := !b

	fmt.Println("Result of AND (a && b):", andResult)
	fmt.Println("Result of OR (a || b):", orResult)
	fmt.Println("Result of NOT (!a):", notA)
	fmt.Println("Result of NOT (!b):", notB)

	fmt.Println("============================")

	x := 5;

	fmt.Println( "X is ",x ,", x > 3 AND x < 10:", x > 3 && x < 10) // true
	fmt.Println( "X is ",x ,", x > 3 OR x < 4:", x > 3 || x < 4) // true
	fmt.Println( "X is ",x ,", NOT (x > 3):", !(x > 3)) // false
}

// Logical Operators are used to determine the logic between variables or values.
/*
	Common logical operators are:
	 --> Are two variables both true?
	 --> Does either of two expressions evaluates to true?
	 --> Is a given expression not true?
*/

/*
	AND Operator (&&):
	 The AND operator returns true if both operands are true.
	 Example: (A && B) is true only if both A and B are true.

	OR Operator (||):
	 The OR operator returns true if at least one of the operands is true.
	 Example: (A || B) is true if either A or B or both are true.

	NOT Operator (!):
	 The NOT operator negates the value of the operand.
	 Example: !A is true if A is false, and false if A is true.
*/
package main;

import (
	"fmt"
)

func main() {
	var a, b int = 12, 5
	fmt.Println("Initial values: a =", a, ", b =", b)

	// Bitwise AND
	andResult := a & b
	fmt.Println("a & b (Bitwise AND):", andResult)

	// Bitwise OR
	orResult := a | b
	fmt.Println("a | b (Bitwise OR):", orResult)

	// Bitwise XOR
	xorResult := a ^ b
	fmt.Println("a ^ b (Bitwise XOR):", xorResult)

	// Left shift
	leftShiftResult := a << 1
	fmt.Println("a << 1 (Left Shift):", leftShiftResult)

	// Right shift
	rightShiftResult := a >> 1
	fmt.Println("a >> 1 (Right Shift):", rightShiftResult)
}

/*
	Bitwise Operators in Go:
	1. &   : Bitwise AND
	2. |   : Bitwise OR
	3. ^   : Bitwise XOR
	4. <<  : Left shift
	5. >>  : Right shift

	These operators are used to perform bit-level operations on integer types in Go.

*/
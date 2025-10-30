package main

import "fmt"

func main() {

	name := "Theja Reddy";
	age := 24;
	height := 5.10;
	weight := 74.5;		// := is the short variable declaration operator
	isStudent := false;
	complexNum := 3 + 4i;
	
	fmt.Println("Name:", name);
	fmt.Println("Age:", age)
	fmt.Printf("Height: %.2f, and Weight: %.2f\n", height, weight); // The formatspecifies only works with Printf.
	fmt.Printf("Is Student: %t\n", isStudent);
	fmt.Printf("Complex Number: %v\n", complexNum);
}

// Data Types In Go

// 1. Strings
/*
	- 16Bytes Memory Allocation 
	- Encodes in double quotes ""
	- Immutable (cannot be changed)
	- Example: "Theja Reddy"
*/

// 2. Integers
/*
	- Whole Numbers (No Decimal Points)
	- Signed & Unsigned
	- Examples: -10, 0, 25, 100
	- for 32 machine defaut is int32 means 4 bytes
	- for 64 machine defaut is int64 means 8 bytes
*/
// 3. Floating-Point Numbers
/*
	- Numbers with Decimal Points
	- Examples: 3.14, -0.001, 2.0
	- float32 (4 bytes) & float64 (8 bytes)
	- foat64 is more precise and recomended to use
*/
// 4. Booleans
/*
	- true or false
	- Example: true, false
	- bool (1 byte)
*/
// 5. Complex Numbers
/*
	- Real & Imaginary Parts
	- Examples: 1+2i, 3-4i
	- complex64 (8 bytes) & complex128 (16 bytes)
	- complex128 is more precise and recommended to use
*/
// 6. Constants
/*
	- Immutable Values
	- Declared with const keyword
	- Example: const pi = 3.14
*/
// 7. Variables
/*
	- Mutable Values
	- Declared with var keyword
	- Example: var name = "Theja"
*/
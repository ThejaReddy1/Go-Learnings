package main;

import ( "fmt" );

var globalVar string = "I am a global"; // package level variable or Global variable

func test() {
	fmt.Println(globalVar);
}

func  main() {
	var name string = "Alice";
	var company string = "Acme Corp";
	var yearsOfExp  float64 = 1.9;
	var designation string = "Senior Software Engineer";
	var isMarried bool = false;
	var NetWorth int = 1000000;

	fmt.Println("Name: ", name);
	fmt.Println(name, "Works at ", company);
	fmt.Printf("%s has %.1f years of experience as a %v\n", name, yearsOfExp, designation);
	fmt.Println("Is Married: ", isMarried);
	fmt.Printf("%v's Net Worth is $%d\n", name, NetWorth);
	test();
}

// variable declaration with var keyword and explicit type
// var variableName dataType = value
// data types: string, int, float64, bool
/*
	variable Scope in Go:
	1. Package Level Scope: Variables declared outside functions are accessible throughout the package.
	2. Function Level Scope: Variables declared within a function are accessible only within that function.
	3. Block Level Scope: Variables declared within a block (e.g., inside loops or conditionals) are accessible only within that block.
*/

// <<==================== ZERO VALUES ==================>>
/*
In Go, when a variable is declared without an explicit initial value, it is assigned a "zero value" based on its data type:

1. Numeric Types (int, float64, etc.): The zero value is 0.
2. String Type: The zero value is an empty string "".
3. Boolean Type: The zero value is false.
4. Pointer Types: The zero value is nil.
5. Composite Types (slices, maps, channels): The zero value is nil.

Example:
	var a int        // a is 0
	var b string     // b is ""
	var c bool       // c is false
	var d *int       // d is nil
	var e []int      // e is nil
	var f float64    // f is 0.0
	pointers, functions, maps, interfaces  is nil by default.
*/


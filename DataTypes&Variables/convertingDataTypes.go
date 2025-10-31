package main;

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 42;
	var b float64 = float64(a); // Convert int to float64
	var c string = string(a);    // Convert int to string (ASCII character)

	fmt.Printf("Integer: %d, Float64: %.2f, String (ASCII): %s\n", a, b, c);

	// Converting string to int
	var strNum string = "123";
	// var strNum string = "123a"; // ==> Uncomment this line to see error handling in action
	var intNum, err = strconv.Atoi(strNum); // Convert string to int
	var intToStr string = strconv.Itoa(a); // Convert int to string (numeric representation)
	fmt.Printf("Error during string to int conversion: %T\n", err); // if convesrsion is successful, err will be nil
	fmt.Printf("String: %s, Converted Int: %d, Int to String: %s\n", strNum, intNum, intToStr);
}

// Data type conversion is called Type Casting.
// In Go, type conversion is done using the syntax: T(v)
// where T is the target type and v is the value to be converted.

// Example:
// var i int = 42
// var f float64 = float64(i)  // Convert int to float64
// var s string = string(i)     // Convert int to string (ASCII character)

// Note: Converting an integer to a string using string(i) converts it to the corresponding ASCII character. 
// To convert an integer to its string representation, use strconv.Itoa(i) from the strconv package.

// Common type conversions:
// 1. int to float64: float64(i)
// 2. float64 to int: int(f)
// 3. int to string (ASCII): string(i)
// 4. string to int: strconv.Atoi(s) (requires error handling)
// 5. int to string (numeric representation): strconv.Itoa(i)

// Always ensure that the conversion is valid to avoid runtime errors.
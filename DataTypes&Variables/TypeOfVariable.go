package main;

import (
	"fmt"
	"reflect"
);

func main() {
	var a int;
	var b string;

	fmt.Println("Type of a is ", reflect.TypeOf(a));
	fmt.Printf("Type of b is %T\n", b);
}

// In go we can determine the type of vaiable in two ways:
// 1. Using %T format specifier with Printf function
// 2. Using reflect.TypeOf() function from reflect package

// Example:
// var age int = 30
// fmt.Printf("Type of age is %T\n", age)  // Using %T format specifier
// fmt.Println("Type of age is ", reflect.TypeOf(age))  // Using reflect.TypeOf() function
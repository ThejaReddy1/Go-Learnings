package main

import "fmt"

// Switch with Conditional Expressions 
// in this case we don't need to mention an expression. After switch we can simply write our conditions in the case blocks itself.
/*
	 Syntax::
		switch {
		case condition1:
			// code to be executed if condition1 is true
		case condition2:
			// code to be executed if condition2 is true
		...
		default:
			// code to be executed if none of the conditions are true
		}
*/
// switch in go impliments implicit break, so we don't need to write break statement at the end of each case block.
func main() {
	var num int = 15

	switch {
		case num < 10:
			fmt.Println("Number is less than 10")
		case num >= 10 && num <= 20:
			fmt.Println("Number is between 10 and 20")
		case num > 20:
			fmt.Println("Number is greater than 20")
		default:
			fmt.Println("This will never be executed")
	}
}
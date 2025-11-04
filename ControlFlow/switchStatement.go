/*
	In Switch case control mechanism, we can execute one block of code among multiple options based on the value of a variable or  an expression.
	The switch statement evaluates an expression, matches the expression's value to a case clause, and executes the associated block of code.

	Syntax::
		switch expression {
		case value1:
			// code to be executed if expression == value1
		case value2:
			// code to be executed if expression == value2
		...
		default:
			// code to be executed if expression doesn't match any case
		}
*/

package main;

import (
	"fmt"
)

func main() {
	var i int = 20

	switch i {
	case 10:
		fmt.Println("Value is 10")
	case 20:
		fmt.Println("Value is 20")
	
	// Multiple values in a single case
	case 2,200:
		fmt.Println("Value is either 2 or 200")
		// fallthrough    // Uncomment this line to see fallthrough in action
	case 1:
		fmt.Println("This will execute because of fallthrough")
	default:
		fmt.Println("Value is neither 10, 20, 200 nor 2")
	}
}

//	Output:	Value is either 2 or 200

// fallthrough
// The fallthrougf keyword is used in swithch case to force execution flow to fall through the successive case block.
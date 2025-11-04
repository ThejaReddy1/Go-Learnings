package main;

import (
	"fmt"
)

func main() {
	var fruit string = "grapes"
	// var fruit string = "apples"  // Uncomment this line to test the other condition
	if fruit == "apples" {
		fmt.Println("Fruit is apple.")
	} else {  // else block need to be in the same line as closing brace of if block otherwise it will throw error
		fmt.Println("Fruit is not apple.")
	}
}

/*
If Conditional in Go:

The if statement is used to execute a block of code based on a condition.
The syntax is as follows:

if condition {
	// code to be executed if condition is true
} else if anotherCondition {
	// code to be executed if anotherCondition is true
} else {
	// code to be executed if none of the above conditions are true
}

Example:

age := 18

if age < 18 {
	fmt.Println("You are a minor.")
} else if age >= 18 && age < 65 {
	fmt.Println("You are an adult.")
} else {
	fmt.Println("You are a senior citizen.")
}

In this example, the program checks the value of 'age' and prints different messages based on the age range.
*/
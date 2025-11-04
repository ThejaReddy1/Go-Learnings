package main;

import (
	"fmt"
)
/*
	Looping with For
	Syntax::
		for initialization; condition; post {
			// code to be executed in each iteration
		}

	Initialization: This step is executed once at the beginning of the loop. It is typically used to initialize a loop counter variable.
	Condition: This is a boolean expression that is evaluated before each iteration of the loop. If the condition evaluates to true, the loop body is executed. If it evaluates to false, the loop terminates.
	Post: This step is executed after each iteration of the loop body. It is typically used to update the loop counter variable.
*/

func main() {
	
	for i := 1; i <=5; i++ {
		fmt.Println("Squares of :", i, "is", i*i);
	}
	// we can skip initialization and post statements if not required
	i := 1;
	for i <= 5 {
		fmt.Println("Cubes of :", i, "is", i*i*i);
		i+=1; // (or i++)
	}

	// if we skip the condition as well, it becomes an infinite loop
	/*
	for {
		fmt.Println("This is an infinite loop");
	}
	*/
}
/*
	In some cases you may want to skip certain iteration based on a condition or exit the loop entirely before it has completed all its iterations. 
	Go provides two control statements, break and continue, to manage the flow of loops.

	break: The break statement is used to exit a loop prematurely. When the break statement is encountered inside a loop, the loop terminates immediately, and the program continues executing the code that follows the loop.

	continue: The continue statement is used to skip the current iteration of a loop and move to the next iteration. When the continue statement is encountered, the remaining code in the current iteration is skipped, and the loop proceeds to the next iteration.
*/

package main;

import (
	"fmt"
)

func main() {
	// break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Print(i, ", ")
	}
	fmt.Println()
	// continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, ", ")
	}
	fmt.Println()
}

package main;

import (
	"fmt"
)

func main() {
	fruit := "grapes"
	// fruit := "apples"  // Uncomment this line to test other conditions
	// fruit := "oranges" // Uncomment this line to test other conditions
	if fruit == "apples" {
		fmt.Println("I love apples.")
	} else if fruit == "oranges" {
		fmt.Println("Oranges are not apples.")
	} else {
		fmt.Println("no appetite for", fruit)
	}
}
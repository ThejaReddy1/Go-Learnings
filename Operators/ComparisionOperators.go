package main;

import (
	"fmt"
)

func main() {
	var city string = "New York"
	var city2 string = "Los Angeles"

	areEqual := city == city2
	areNotEqual := city != city2
	isGreater := city > city2 // it returns true if city comes after city2 lexicographically
	isLess := city < city2 

	fmt.Println("Are cities equal?", areEqual)
	fmt.Println("Are cities not equal?", areNotEqual)
	fmt.Println("Is city greater than city2?", isGreater)
	fmt.Println("Is city less than city2?", isLess)

	a := 10
	b := 20

	isAGreater := a > b
	isALess := a < b
	isAEqualB := a == b
	isANotEqualB := a != b

	fmt.Println("Is a greater than b?", isAGreater)
	fmt.Println("Is a less than b?", isALess)
	fmt.Println("Is a equal to b?", isAEqualB)
	fmt.Println("Is a not equal to b?", isANotEqualB)
}
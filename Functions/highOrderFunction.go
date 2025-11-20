package main

import (
	"fmt"
	"math"
)
/*
	High Order Functions:
	High Order Function is a function that receives a function as an argument or return a unction as output.
*/
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func circleParimeter(r float64) float64 {
	return 2 * math.Pi * r
}

func circleDiameter(r float64) float64 {
	return 2 * r
}
func printResult(radius float64, calcFunc func(r float64) float64) {
	result := calcFunc(radius)
	fmt.Printf("Result: %.2f\n", result)
	fmt.Println("Thank you!")
}
func getFunction(query int) func(r float64) float64 {
	query_to_func := map[int]func(r float64) float64{
		1: circleArea,
		2: circleParimeter,
		3: circleDiameter,
	}
	return query_to_func[query]
}
func main() {
	test := func() {
		fmt.Println("Hello!")
	}
	test() // Calling the function stored in variable 'test'

	func() {
		fmt.Println("Hello from test2!")
	}() // Immediately Invoked Function Expression (IIFE)

	var query int
	var radius float64

	fmt.Print("Enter the radius of the circle: ")
	fmt.Scanf("%f", &radius)
	fmt.Print("Choose an option :- \n1: Area\n 2: Perimeter\n 3: Diameter\n ")
	fmt.Scanf("%d", &query)

	printResult(radius, getFunction(query))
	

}
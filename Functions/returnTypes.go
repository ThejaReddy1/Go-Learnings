package main;

import (
	"fmt"
)
/*
	If the function returns a value with different data type, than you specified in the return type, it will result in a compile-time error. `indicating a type mismatch.`

	Function can return multiple values by specifying multiple return types in the function signature and returning corresponding values in the return statement.
*/

func add(x int, y int) int {
	return x + y;
}

func muldiv (x int, y int) (int,int) {
	return x * y, x / y;
}

/*
	Named return values:
	Go allows giving the names to the rturn or result parameters of a functions, in the function signature or definition. This is known as named return values.
	When using named return values, you can omit the variable declaration in the return statement, as the named return values are already defined in the function signature.
*/
func operations(x int, y int) (sum int, diff int) {
	sum = x + y
	diff = x - y 
	return
}

func main() {
	a := 10
	b := 20
	sum := add(a,b)
	fmt.Println("Sum of ", a, " and ", b, " is: ", sum)
	multiple,division := muldiv(a,b)
	fmt.Println("Multiplication of ", a, " and ", b, " is: ", multiple)
	fmt.Println("Division of ", a, " and ", b, " is: ", division)
	sum2,subtract := operations(a,b)

	fmt.Println("Using Named Return Values - Sum: ", sum2, " Subtraction: ", subtract)

	
}
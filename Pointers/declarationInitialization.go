package	main

import "fmt"
/*
	Declaring a pointer
		var <pointer_name> *<data_type>
		ex: var ptr *int
			var strPtr *string
	Initializing a pointer
		var <pointer_name> *<data_type> = &<variable_name>
		ex: var ptr *int = &i		// ptr stores the address of i
			var strPtr *string = &str	// strPtr stores the address of str
*/
func main() {
	var i *int
	var strPtr *string

	x := 100 
	y := "Joe"

	i = &x			// i now holds the address of x
	strPtr = &y		// strPtr now holds the address of y

	fmt.Println("Address of i", i,"Value of i:", *i) // Output: Value of i: <nil>
	fmt.Println("Address of strPtr", strPtr,"Value of strPtr:", *strPtr)
}
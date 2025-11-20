package main

import (
	"fmt"
)

/*
	Defer Statement:
	The `defer` statement delays the execution of a function until the surrounding function returns.
	Deferred functions are executed in LIFO (Last In, First Out) order just before the surrounding function returns.
	Defer is commonly used to close files, release resources, or unlock mutexes.
*/
func printName(name string) {
	fmt.Println("name: ", name)
}
func printRollNo(rollno int) {
	fmt.Println("roll no: ", rollno)
}
func printAddress(address string) {
	fmt.Println("address: ", address)
}

func main() {
	printName("John Doe")
	defer printRollNo(101)
	defer printAddress("123 Main St, Cityville")
	fmt.Println("End of main function")
}
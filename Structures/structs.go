package main

import (
	"fmt"
)

/*
	-> struct is a user defined data type
	-> a structure that groups together data elements of different types
	-> providing a way to reference a series of gouped values through a single variable name.
	-> used when it makes sense to group or associate two or more data variables.

	syntax:
		type <struct_name> struct {
			<variable_name> <data_type>
			<variable_name> <data_type>
			<variable_name> <data_type>
		}
*/

func main() {
	type person struct {
		name string
		age int
		employee bool
	}
	// creating a variable of type person
	var p1 person
	// assigning values to the variables
	p1.name = "James"
	p1.age = 21
	p1.employee = true

	fmt.Println(p1)
	fmt.Printf("%+v\n", p1)
	fmt.Println(p1.name)

	// Initializing Structs

	var p2 *person = new(person)
	fmt.Printf("%+v\n", p2)
	p2.age = 55
	// or
	p3 := person{"James", 21, true}
	fmt.Printf("%+v\n", p3)
	// or
	p4 := person{
		name: "Marry",
		age: 21,
		employee: false,
	}
	fmt.Printf("%+v\n", p4)

	// Accessing Structs
	fmt.Println(p4.name)
	fmt.Println(p2.age)
}
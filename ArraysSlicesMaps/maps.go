package main;

import (
	"fmt"
);

/*
	Maps in Go:
	` A map is an unordered collection of key-value pairs.`
	Maps are used to store data in key-value pairs, where each key is unique and maps to a specific value.
		Dictonary in Python
		HashMap in Java
		Object in JavaScript (used as a map)

	Declaring a map:
		var <map_name> map [<key_data_type>]<value_data_type>	// Nil map
	Or
		<map_name> := map[<key_data_type>]<value_data_type>{"key":"value", "key":"value", ...} // Initialized map
	Or
	   <map_name> := make(map[<key_data_type>]<value_data_type>)
	Or
	   <map_name> := map[<key_data_type>]<value_data_type>{<key1>:<value1>, <key2>:<value2>, ...}

	Accessing values in a map:
	   value := <map_name>[key]

	Adding or updating values in a map:
	   <map_name>[key] = value

	Deleting a key-value pair from a map:
	   delete(<map_name>, key)

	Checking if a key exists in a map:
	   value, exists := <map_name>[key]
	   if exists {
	       // key exists
	   } else {
	       // key does not exist
	   }
*/

func main() {
	// Declaring and initializing a map

	var mapNill map[string]int // Nil map
	fmt.Println("Nil map: ", mapNill)

	// Using make function to create a map
	mapMake := make(map[string] int)
	fmt.Println("Map created using make: ", mapMake)

	// Initialized map
	mapInitialized := map[string] int {"Alice": 25, "Bob": 30, "Charlie": 35}
	fmt.Println("Initialized map: ", mapInitialized) 

	// Accessing values in a map
	ageOfAlice := mapInitialized["Alice"]
	fmt.Println("Age of Alice: ", ageOfAlice)

	// Adding or updating values in a map
	mapInitialized["David"] = 40 // Adding new key-value pair
	fmt.Println("After adding David: ", mapInitialized)

	mapInitialized["Alice"] = 26 // Updating value for existing key
	fmt.Println("After updating Alice's age: ", mapInitialized)

	// Deleting a key-value pair from a map
	delete(mapInitialized, "Alice")
	fmt.Println("After deleteing Alice: ", mapInitialized)

	// Checkiing if a key exists in a map

	age, exists := mapInitialized["Bob"]	 // returns value and boolean
	if exists {
		fmt.Println("Bob's age is: ", age)
	} else {
		fmt.Println("Bob does not exist in the map")
	}

	// Iterating over a map
	for key, value := range mapInitialized {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// Getting the length of a map
	lengthOfMap := len(mapInitialized)
	fmt.Println("Length of the map: ", lengthOfMap)

	// Note: Maps in Go are unordered, so the order of key-value pairs may vary when iterating over them

	// Truncating a map
		mapInitialized = make(map[string]int) // re-initializing the map to an empty map
		fmt.Println("After truncating the map: ", mapInitialized)
	// Or simply delete all keys using a loop
		for key := range mapInitialized {
			delete(mapInitialized, key)
		}
} 
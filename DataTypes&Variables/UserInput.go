package main;
import (
	"fmt"
)

func scanftest() {
	var a string;
	var b int;
	fmt.Print("Enter a string and an integer: ");
	count, err := fmt.Scanf("%s %d", &a, &b);

	fmt.Printf("You entered: %s and %d\n", a, b);
	fmt.Println("Number of items successfully scanned: ", count);
	fmt.Println("Error: ", err);
}

func main() {
	var name string;
	fmt.Print("Enter your name: ");
	fmt.Scanf("%s", &name);  // Scanf fmt.Sacnf("%<format_specifier>(s)", &variable (Or Object_arguments) )

	fmt.Println("Hello, ", name);

	//multiple inputs
	var designation string;
	var yearsOfExp float64;

	fmt.Print("Enter your designation and years of experience: ");
	fmt.Scanf("%s %f", & designation, &yearsOfExp);

	fmt.Printf("%s has %.1f years of experience as a %s\n", name, yearsOfExp, designation);

	scanftest();
}

// fmt.Scanf function Returns two values:
// 1. number of items successfully scanned (count)
// 2. error value (nil if no error occurred)
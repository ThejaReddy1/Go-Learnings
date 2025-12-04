package main
import "fmt"
/*
	Meth
*/

type Car struct {
	brand string
	accelerate int
}

// Method takes Photocopy of car struct and just utilize it any changes to the data will not affect the original.
func (c Car) drive() {
	fmt.Println(c.brand, "is going at Speed, ",c.accelerate)
}

// Method takes Reference of Original struct and any modifications to the data will reflect in the actual original data 

func (c *Car) accelerateCar(speed int) {
	c.accelerate += speed
}

func main() {
	car := Car{ brand: "BMW", accelerate: 0 }

	car.drive()

	car.accelerateCar(20)

	car.drive()

	fmt.Println(car.accelerate)
}

/*
	Methods are just functions with a receiver argument.
	Methods are just like functions but they are defined with a receiver.
			Syntax: 
				func (receiverName receiverType) methodname() {
					// Method body
				}
					( or )
				func (receiverName receiverType) methodName(parameterList) (returnTypeList) {
					// Method body
				}
	Just like functions/Methods in java Class, you define data(fields) and behavior (methods) inside a class block.
	You define data in a struct and attach behavior (methods) to it separately. But logically, they form the same "Object".
	
*/



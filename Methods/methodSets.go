package main
import "fmt"
/*
	Method Sets are nothing but a set of methods that are available with same data type and it is a useful way to encapsulate functionlity.
*/

type Student struct {
	name string
	grades []int
}

func (s *Student) studentName() {
	fmt.Println("Student Name: ",s.name)
	fmt.Println("Student Grades: ", s.grades)
}

func (s *Student) percentage() float32 {
	sum := 0
	for _,v := range s.grades {
		sum += v
	}
	percentage := (float32(sum)/ float32(len(s.grades) * 100)) * 100

	return percentage
}

func main() {
	// s1 := Student{"John", []int{70,90,80,85}}
	s1 := Student{name: "john", grades: []int{70,90,80,85}}
	s1.studentName()
	fmt.Printf("Student %s has %.2f%% percentage\n",s1.name, s1.percentage())
}
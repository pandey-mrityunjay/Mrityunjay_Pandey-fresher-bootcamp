package main

import "fmt"

//Structs: is a collection of fields taht define a custom type
//Defines real world entities

type Person struct {
	Name string
	Age  int16
}

func main() {
	p := Person{Name: "John", Age: 25}
	fmt.Println(p.Name, p.Age)
	p.Birthday()
	fmt.Println(p.Name, p.Age)

	p.Greet()
	p.SetAge(30)
	fmt.Println(p.Age)
}

func (p *Person) Birthday() {
	p.Age++
}

//when to use structs:
//1: when defining real world objects
//2: when organizing data logically

/*
	1.Use pointers when modifying data efficiently
	2.Use arrays when the size is fixed.
	3.Use slices when the size is dynamic
	4.Use maps when storing key-value relationships
	5.Use range when iterating over collections
	6.use structs when modelling real-world entities

*/

//A method in Go is a function with a reciever that is associated with astruct type.
//Method allows us to define behaviour specific to that struct

/*
Syntax
func( reciever type)
Method name (parameters) returnType{

}
*/

func (p Person) Greet() {
	fmt.Println("Hello,my name is", p.Name)
}

func (p *Person) SetAge(newAge int) {
	p.Age = int16(newAge)
}

// Pointer is a variable that stores the memory address of the another variable instead of storing a value directly. It holds the location where the value is stored
// Use Cases:
// 1: Pass by reference instead of pass by value
// 2: Efficient memory management
// 3: Works with structs and large data efficiently
package main

import "fmt"

func modify(x *int) {
	*x = *x * 2

}
func main() {
	num := 10
	fmt.Println("Before:", num)

	modify(&num)
	fmt.Println("After:", num)
}

// &num gets the memory address of num
// *x is the deferencing operator that access/modifies the value at the memory location

package main

import "fmt"

func main(){
	for i:=0;i<=5;i++{
		fmt.Println("Count:",i)
	}

	for_as_while()
	if_case()
	switch_case()
	defer_usage()
}

//for loop as while 
func for_as_while(){
	i:=0
	for i<5{
	fmt.Println("Value of i:",i)
	i++
	}
}

func if_case(){
	age:=18

	if age>=18{
		fmt.Println("You're an adult")
	}	else {
		fmt.Println("You're a minor")
	}
}

func switch_case(){
	day:=3

	switch day{
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	default:
		fmt.Println("Sunday")
	}

}



//DEFER KEYWORD 
// defer fmt.println("WORLD") executes after the main function completes 
func defer_usage(){
	defer fmt.Println("WORLD")
	fmt.Println("HELLO")
}


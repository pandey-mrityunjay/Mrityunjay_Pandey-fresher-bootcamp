package main

import (
	"fmt"
	"math"
)

const inflateRate = 3.5

func main() {
	var investmentAmount, expectedReturnRate float64
	var years int64

	fmt.Print("Investment Amount:")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate:")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years:")
	fmt.Scan(&years)

	fmt.Println("")

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))
	futurRealValue := futureValue / math.Pow(1+inflateRate/100, float64(years))

	fmt.Print(futurRealValue)
	fv1, fv2 := calculateFutureValues(10000, 4.5, 5)
	fmt.Printf("%v \n %v", fv1, fv2)
	outputText("futurRealValue")
}

// function declarations
func outputText(text string) {
	fmt.Print(text)
}

// functions: return values and variable scopes
func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years int) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))
	fv1 := fv / math.Pow(1+inflateRate/100, float64(years))
	return fv, fv1
}

//Variadic function: A variadic function is a function that takes a variable number of arguments, In Go, you define a varidic function
// by using ... nefore the parameter type

func variadic_func(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

//passing a slice to a variadic function
// nums := []int{1,2,3,4,5,6}
//fmt.Println(sum(nums...))

//CLOSURES: is a function that references variables from outside its body

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// counter returns an anonymous function that increments count
// since count is defined in counter, the inner function remembers its value even after counter has finished execution
func closure_Call() {
	increment := counter()

	fmt.Println(increment()) //Output:1
	fmt.Println(increment()) //Output:2
	fmt.Println(increment()) //Output:3
}

//Recursion
//When a function calls itself, is known as a recursive function

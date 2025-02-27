package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000 // type explicit assignment float 64
	// var investmentAmount , years float64:= 1000,10       =>   multiple declarations in a single line 
	// all the variables should be of same type
	// investmentAmount,years:=1000.0,10.0

	//Alternative Type Declaration
	//where the type is to be inferred by the go compiler use := this assignment operator
	expectedReturnRate := 5.5

	
	var years float64 = 10

	//user input syntax
	//fmt.Scan(&investmentAmount)

	var futureValue = (investmentAmount) * math.Pow(1+expectedReturnRate/100, (years))

	fmt.Println(futureValue)

}

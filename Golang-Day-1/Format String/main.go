package main

import (
	"fmt"
	"math"
)

func main() {
	const inflateRate = 2.5
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

	//fmt.Println("Future Vale",futureValue)
	//adding information to the output

	//1:printf to print with formatting
	fmt.Printf("Future Value: %v\n Future Value Adjusted: %v \n", futureValue, futurRealValue)

	//2:formatting floats in strings
	// %.nf  n is a number that defines number of decimal places
	fmt.Printf("Future Value: %.0f\n Future Value Adjusted: %.0f \n", futureValue, futurRealValue)

	//3; storing the formatted value to a variable using sprintf
	formattedString:=fmt.Sprintf("Future Value: %.0f\n Future Value Adjusted: %.0f \n", futureValue, futurRealValue)

	fmt.Println(formattedString)

	//4:building a multiline strings
	

	//fmt.Println(futurRealValue)
}

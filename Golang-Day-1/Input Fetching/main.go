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

	fmt.Println("Future Vale",futureValue)
	fmt.Println(futurRealValue)
}
/*The fmt.Scan() function is a great function for easily fetching & using user input through the command line.

But this function also has an important limitation: You can't (easily) fetch multi-word input values. 
Fetching text that consists of more than a single word is tricky with this function. */
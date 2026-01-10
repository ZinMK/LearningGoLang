package main

import (
	"fmt"
	"math"
)

func main() {	

	const inflation = 2.5
	var years float64
	var investmentAmount, expectedReturnRate float64

	fmt.Print("Enter Investment Ammount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter years: ")
	fmt.Scan(&years)
	fmt.Print("Enter expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow((1+ expectedReturnRate/100),years)
	futureRealValue := futureValue / math.Pow(1+inflation/100, years)
	
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
	
}

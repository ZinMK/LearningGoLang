package main

import (
	"fmt"
)

func main(){
	
	var revenue, expenses, taxRate float64
	fmt.Println("Enter your Profit: ")
	fmt.Scan(&revenue)
	fmt.Println("Enter Your Expenses: ")
	fmt.Scan(&expenses)
	fmt.Println("Enter your TaxRate: ")
	fmt.Scan(&taxRate)

	var earnings_before_tax float64 = revenue - expenses
	var earnnings_after_tax float64 = earnings_before_tax * (1 - (taxRate/100))
	var ratio float64 = earnings_before_tax/earnnings_after_tax

	fmt.Print("Earnings Before Tax: ")
	fmt.Println(earnings_before_tax)

	fmt.Print("\nEarnings After Tax: ")
	fmt.Println(earnnings_after_tax)

	fmt.Print("Ratio")
	fmt.Println(ratio)




}
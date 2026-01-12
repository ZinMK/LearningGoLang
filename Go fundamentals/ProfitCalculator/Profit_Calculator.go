package main

import (
	"fmt"
	"os"
)

func write_details_to_file(revenue, profit, ratio float64) {
	results := fmt.Sprint(revenue, profit, ratio)
	os.WriteFile("numbers.txt", []byte(results), 0644)
	
}

func main(){
	
	var revenue, expenses, taxRate float64
	
	get_user_input("Enter your revenue: ", &revenue)
	get_user_input("Enter your expenses: ", &expenses)
	get_user_input("Enter your TaxRate: ", &taxRate)
	
	
	earnings_before_tax , earnnings_after_tax, ratio := calculations(revenue, expenses, taxRate)

	write_details_to_file(earnings_before_tax,earnnings_after_tax,ratio)

	formattedEBT := fmt.Sprintf("Earnings Before Tax: %.2f\n", earnings_before_tax)
	formattedEAT := fmt.Sprintf("Earnings After Tax: %.2f\n",earnnings_after_tax)
	formattedRatio := fmt.Sprintf("Earning Ratio %.2f", ratio)
	

	fmt.Println(formattedEAT,formattedEBT,formattedRatio)
 
}

func get_user_input(outputStr string,assignmentVariable *float64) {
	fmt.Println(outputStr)
	fmt.Scan(assignmentVariable)
	enteredvalue := *assignmentVariable
	if enteredvalue == 0 || enteredvalue<0 {
		panic("You entered an invalid value now the program has crashed")
	}

}	

func calculations(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	EBT := revenue - expenses
	EAT :=  EBT * (1 - (taxRate/100))
	Ratio := EBT/EAT
	return EBT, EAT, Ratio
}


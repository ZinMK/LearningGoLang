package main

import (
	"fmt"
)

func main(){
	
	var revenue, expenses, taxRate float64
	
	get_user_input("Enter your revenue: ", &revenue)
	get_user_input("Enter your expenses: ", &expenses)
	get_user_input("Enter your TaxRate: ", &taxRate)
	
	earnings_before_tax , earnnings_after_tax, ratio := calculations(revenue, expenses, taxRate)

	formattedEBT := fmt.Sprintf("Earnings Before Tax: %.2f\n", earnings_before_tax)
	formattedEAT := fmt.Sprintf("Earnings After Tax: %.2f\n",earnnings_after_tax)
	formattedRatio := fmt.Sprintf("Earning Ratio %.2f", ratio)
	
	fmt.Println(formattedEAT,formattedEBT,formattedRatio)
 
	

}

func get_user_input(outputStr string,assignmentVariable *float64) {
	fmt.Println(outputStr)
	fmt.Scan(assignmentVariable)

}	

func calculations(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	EBT := revenue - expenses
	EAT :=  EBT * (1 - (taxRate/100))
	Ratio := EBT/EAT
	return EBT, EAT, Ratio
}

